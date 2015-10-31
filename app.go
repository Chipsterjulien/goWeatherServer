package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/itsjamie/gin-cors"
	"strconv"
	"time"
)

// Struct to get only temperature. time is useless
type OnlyTemperature struct {
	Temp float64
}

// Struct permit to have DB access on method
type Ressource struct {
	db *gorm.DB
}

// Init DB
func NewRessource(db *gorm.DB) Ressource {
	return Ressource{
		db: db,
	}
}

// Remove temperature with id
func (r *Ressource) DeleteTemperature(c *gin.Context) {
	log := logging.MustGetLogger("log")
	id := c.Params.ByName("id")

	log.Debug("id for delete a temperature is \"%s\"", id)

	if id != "" {
		var temperature Temperature
		r.db.First(&temperature, id)
		log.Debug("Temperature to delete is:")
		log.Debug("  temperature: %f\n  date: %v", temperature.Temp, temperature.Date)
		r.db.Delete(&temperature)

		c.JSON(200, gin.H{"id #" + id: "delete"})
	} else {
		c.JSON(404, gin.H{"error": "Unable to remove temperature. No id given"})
		log.Warning("Unable to remove temperature. No id given")
	}
}

// Get all temperatures in database
func (r *Ressource) GetTemperatures(c *gin.Context) {
	log := logging.MustGetLogger("log")
	temperatures := []Temperature{}

	r.db.Find(&temperatures)

	if viper.GetString("logtype") == "debug" {
		log.Debug("Temperatures in DB are:")
		for _, temp := range temperatures {
			log.Debug("  Temperature: %f\n  Date: %v", temp.Temp, temp.Date)
		}
	}

	if len(temperatures) == 0 {
		c.JSON(404, gin.H{"error": "No temperature in database"})
	} else {
		c.JSON(200, temperatures)
	}
}

// Get all temperatures in database, filter and keep the temperature of the day
func (r *Ressource) GetTemperaturesOfTheDay(c *gin.Context) {
	temperatures := []Temperature{}
	temperaturesWillBeReturned := []Temperature{}

	r.db.Find(&temperatures)
	if len(temperatures) == 0 {
		c.JSON(404, gin.H{"error": "No temperature in database"})
	} else {
		now := time.Now()
		for _, t := range temperatures {
			if t.Date.Day() == now.Day() && t.Date.Month() == now.Month() && t.Date.Year() == now.Year() {
				temperaturesWillBeReturned = append(temperaturesWillBeReturned, t)
			}
		}
		if len(temperaturesWillBeReturned) == 0 {
			c.JSON(404, gin.H{"error": "No temperature in database"})
		} else {
			c.JSON(200, temperaturesWillBeReturned)
		}
	}
}

// Get all temperatures in database, filter and keep the temperature of the month
func (r *Ressource) GetTemperaturesOfTheMonth(c *gin.Context) {
	temperatures := []Temperature{}
	temperaturesWillBeReturned := []Temperature{}

	r.db.Find(&temperatures)
	if len(temperatures) == 0 {
		c.JSON(404, gin.H{"error": "No temperature in database"})
	} else {
		now := time.Now()
		for _, t := range temperatures {
			if t.Date.Month() == now.Month() && t.Date.Year() == now.Year() {
				temperaturesWillBeReturned = append(temperaturesWillBeReturned, t)
			}
		}
		if len(temperaturesWillBeReturned) == 0 {
			c.JSON(404, gin.H{"error": "No temperature in database"})
		} else {
			c.JSON(200, temperaturesWillBeReturned)
		}
	}
}

// Get all temperatures in database, filter and keep the temperature of the year
func (r *Ressource) GetTemperaturesOfTheYear(c *gin.Context) {
	temperatures := []Temperature{}
	temperaturesWillBeReturned := []Temperature{}

	r.db.Find(&temperatures)
	if len(temperatures) == 0 {
		c.JSON(404, gin.H{"error": "No temperature in database"})
	} else {
		now := time.Now()
		for _, t := range temperatures {
			if t.Date.Year() == now.Year() {
				temperaturesWillBeReturned = append(temperaturesWillBeReturned, t)
			}
		}
		if len(temperaturesWillBeReturned) == 0 {
			c.JSON(404, gin.H{"error": "No temperature in database"})
		} else {
			c.JSON(200, temperaturesWillBeReturned)
		}
	}
}

// Post a temperature into database
func (r *Ressource) PostTemperature(c *gin.Context) {
	log := logging.MustGetLogger("log")
	var temp OnlyTemperature

	c.Bind(&temp)

	log.Debug("Temperature: %f", temp.Temp)

	temperature := Temperature{
		Temp: temp.Temp,
		Date: time.Now(),
	}

	r.db.Save(&temperature)
	c.JSON(200, temperature)
}

// Main function
func startApp(db *gorm.DB) {
	log := logging.MustGetLogger("log")

	if viper.GetString("logtype") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.Default()
	r := NewRessource(db)
	
	g.Use(cors.Middleware(cors.Config {
	    Origins:        "*",
	    Methods:        "GET, PUT, POST, DELETE",
	    RequestHeaders: "Origin, Authorization, Content-Type",
	    ExposedHeaders: "",
	    MaxAge: 50 * time.Second,
	    Credentials: true,
	    ValidateHeaders: false,
	}))

	v1 := g.Group("api/v1")
	{
		v1.GET("/temperatures", r.GetTemperatures)
		//v1.GET("/temperatures_of_the_day", r.GetTemperaturesOfTheDay)
		//v1.GET("/temperatures_of_the_month", r.GetTemperaturesOfTheMonth)
		//v1.GET("/temperatures_of_the_year", r.GetTemperaturesOfTheYear)
		v1.POST("/temperature", r.PostTemperature)
	}
	log.Debug("Port: %d", viper.GetInt("server.port"))
	g.Run(":" + strconv.Itoa(viper.GetInt("server.port")))
}

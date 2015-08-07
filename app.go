package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type OnlyTemperature struct {
	Temp float64
}

type Ressource struct {
	db *gorm.DB
}

func NewRessource(db *gorm.DB) Ressource {
	return Ressource{
		db: db,
	}
}

// Remove temperature with id
func (r *Ressource) DeleteTemperature(c *gin.Context) {
	log := logging.MustGetLogger("log")
	id := c.Params.ByName("id")

	if id != "" {
		var temperature Temperature
		r.db.First(&temperature, id)
		r.db.Delete(&temperature)

		c.JSON(200, gin.H{"id #" + id: "delete"})
	} else {
		c.JSON(404, gin.H{"error": "Unable to remove temperature. No id given"})
		log.Warning("Unable to remove temperature. No id given")
	}
}

// Get all temperatures in database
func (r *Ressource) GetTemperature(c *gin.Context) {
	temperatures := []Temperature{}

	r.db.Find(&temperatures)
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
	var temp OnlyTemperature

	c.Bind(&temp)
	temperature := Temperature{
		Temp: temp.Temp,
		Date: time.Now(),
	}

	r.db.Save(&temperature)
	c.JSON(200, temperature)
}

func startApp(db *gorm.DB) {
	g := gin.Default()
	r := NewRessource(db)

	v1 := g.Group("api/v1")
	{
		v1.GET("/temperatures", r.GetTemperature)
		v1.GET("/temperatures_of_the_day", r.GetTemperaturesOfTheDay)
		v1.GET("/temperatures_of_the_month", r.GetTemperaturesOfTheMonth)
		v1.GET("/temperatures_of_the_year", r.GetTemperaturesOfTheYear)
		v1.POST("/temperature", r.PostTemperature)
	}
	g.Run(":" + strconv.Itoa(viper.GetInt("port")))
}

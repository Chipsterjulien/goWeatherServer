package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
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

// Post a temperature into database
func (r *Ressource) PostTemperature(c *gin.Context) {
	//log := logging.MustGetLogger("log")
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
		v1.POST("/temperature", r.PostTemperature)
	}
	g.Run(":8080")
}

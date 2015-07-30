package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func startApp(db *gorm.DB) {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/temperatures", GetTemperatures)
	}
	r.Run(":8080")
}

func GetTemperatures(c *gin.Context) {
	temperatures := []Temperature{}
	log.Fatal(c)

	_ = temperatures
}

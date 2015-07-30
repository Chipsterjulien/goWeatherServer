package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/op/go-logging"
	"os"
	"time"
)

type Temperature struct {
	Id   int `gorm:primary_key`
	date string
	temp time.Time
}

func Initdb() *gorm.DB {
	log := logging.MustGetLogger("log")
	db, err := gorm.Open("sqlite3", "temperature.db")
	if err != nil {
		log.Critical("Unable to open db file:", err)
		os.Exit(1)
	}
	db.LogMode(true)
	db.CreateTable(new(Temperature))
	db.DB().Ping()

	return &db
}

package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	confFilename := "goweather"
	confPath := "."
	logFilename := "errors.log"

	//confFilePath := "/etc/goweather"
	//dbPath := "/var/lib/goweather"
	//logFilename := "/var/log/goweather/errors.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)

	fmt.Println(viper.Get("database"))
	fmt.Println(viper.GetString("database.filename"))
	fmt.Println(viper.IsSet("database.path"))
	os.Exit(0)

	dbmap := Initdb()
	startApp(dbmap)
}

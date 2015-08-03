package main

func main() {
	confFilename := "goweather"
	confPath := "."
	logFilename := "errors.log"

	//confPath := "/etc/goweather"
	//logFilename := "/var/log/goweather/errors.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)

	dbmap := Initdb()
	startApp(dbmap)
}

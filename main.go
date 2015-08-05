package main

func main() {
	confPath := "/etc/goWeatherServer"
	confFilename := "goWeatherServer"
	logFilename := "/var/log/goWeatherServer/errors.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)

	dbmap := Initdb()
	startApp(dbmap)
}

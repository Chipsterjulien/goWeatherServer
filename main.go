package main

func main() {
	confPath := "/etc/goweatherserver"
	confFilename := "goweatherserver"
	logFilename := "/var/log/goweatherserver/errors.log"

	fd := initLogging(&logFilename)
	defer fd.Close()

	loadConfig(&confPath, &confFilename)

	dbmap := Initdb()
	startApp(dbmap)
}

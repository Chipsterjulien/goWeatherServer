package main

func main() {
	fd := initLogging()
	defer fd.Close()

	dbmap := Initdb()
	startApp(dbmap)
}

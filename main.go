package main

import (
	"api_sensor/config"
	"api_sensor/core"
	"api_sensor/router"
)

func main() {
	config.Dotenv()
	config.DBconnect()
	config.Migrate()

	go core.Main()

	router.Main()
}

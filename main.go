package main

import (
	"km-kelas-e/config"
	"km-kelas-e/migrate"
	"km-kelas-e/routes"
)

func main() {
	//initiateDB
	config.InitDB()
	migrate.AutoMigrate()

	//initRoutes
	e := routes.New()

	//starting the server
	e.Start(":1234")
}

package application

import (
	"Cih2001/GoTruck/web/controllers"
)

func setupRoutes() {
	setupHome()
}

func setupHome() {
	Server.GET("/", controllers.HomeIndex)
}

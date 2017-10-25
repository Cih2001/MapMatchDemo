package application

import (
	"Cih2001/MapMatchDemo/controllers"
)

func setupRoutes() {
	setupHome()
}

func setupHome() {
	Server.GET("/", controllers.HomeIndex)
}

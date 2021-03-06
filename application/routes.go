package application

import (
	"Cih2001/MapMatchDemo/controllers"
)

func setupRoutes() {
	setupHome()
	setupApis()
}

func setupHome() {
	Server.GET("/", controllers.HomeIndex)
}

func setupApis() {
	Server.POST("/api/add-point", controllers.AddPoint)
	Server.POST("/api/get-projections", controllers.GetProjections)
	Server.POST("/api/get-roads", controllers.GetRoads)	
	Server.GET("/api/show-default-path", controllers.Test)
}

package controllers

import (
	"Cih2001/MapMatchDemo/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	path := new(models.Path)
	path = path.Find("")

	return c.JSON(http.StatusOK, path)
}
func AddPoint(c echo.Context) error {
	form := new(AddPointForm)
	if err := c.Bind(form); err != nil {
		log.Println("error binding", err)
	}
	path := new(models.Path)
	path = path.Find(form.PathName)
	path.AddPoint(form.Lat, form.Long)
	return c.JSON(http.StatusOK, map[string]string{"msg": "ok"})
}

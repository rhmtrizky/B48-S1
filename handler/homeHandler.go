package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)



func home(c echo.Context) error {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}
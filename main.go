package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	
	route := echo.New()

	route.Static("/assets", "assets")



	route.GET("/home", home)
	route.GET("/contact", contact)
	route.GET("/project", project)
	route.GET("/testimonial", testimonial)
	route.GET("/detail-project/:id", detailProject)
	route.POST("/add-project", addProject)

	route.Logger.Fatal(route.Start("localhost:5000"))
}

func home(c echo.Context) error {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	temp, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}

func project(c echo.Context) error {
	temp, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}
func testimonial(c echo.Context) error {
	temp, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}

func detailProject(c echo.Context) error {
	id := c.Param("id")

	temp, err := template.ParseFiles(("views/project-detail1.html"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	detailProject := map[string]interface{}{
		"id" : id,
		"title" : "Education Mobile App",
		"startDate": "18 July 2021",
		"endDate": "20 Maret 2023",
		"duration": "2 Years",
	}
	return temp.Execute(c.Response(), detailProject)
}

func addProject(c echo.Context) error {

	projectName := c.FormValue("nameProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("desc")
	nodejs := c.FormValue("node")
	reactjs := c.FormValue("react")
	javaScript := c.FormValue("javaScript")
	golang := c.FormValue("go")
	image := c.FormValue("image")


	fmt.Println("Project Name: " , projectName)
	fmt.Println("Start Date: " , startDate)
	fmt.Println("End Date: " , endDate)
	fmt.Println("Description: " , description)
	fmt.Println("Image: " , image)


	if nodejs == "checked" {
		fmt.Println("Node js: ✅")
	} else {
		fmt.Println("Node js:❌")
	}
	
	if reactjs == "checked" {
		fmt.Println("React js: ✅")
	} else {
		fmt.Println("react js:❌")
	}

	if javaScript == "checked" {
		fmt.Println("JavaScript: ✅")
	} else {
		fmt.Println("JavaScript: ❌")
	}
	if golang == "checked" {
		fmt.Println("Golang: ✅")
	} else {
		fmt.Println("Golang: ❌")
	}

	return c.Redirect(http.StatusMovedPermanently, "/project")

}
package main

import (
	"context"
	"fmt"
	"html/template"
	"my-web-module/connection"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id			int
	ProjectName string
	StartDate time.Time
	EndDate time.Time
	Duration string
	Description string
	Nodejs bool
	Reactjs bool
	JavaScript bool
	Golang bool
	Image  string
}

// var dataProjects = []Project{
// 	{
// 		Id: 1,
// 		ProjectName : "Education Mobile App",
// 		StartDate : time.Now(),
// 		EndDate : time.Now(),
// 		Duration : "3 minggu",
// 		Description : "Apalikasi ini dibuat dri tahun kmren yang diginakan untuk belajar online",
// 		Nodejs: true,	
// 		Reactjs: true,
// 		JavaScript: true,
// 		Golang: false,
// 	},
// 	{
// 		Id: 2,
// 		ProjectName : "Music Mobile App",
// 		StartDate : time.Now(),
// 		EndDate : time.Now(),
// 		Duration : "4 minggu 2 hari",
// 		Description : "Apalikasi ini dibuat dri tahun kmren yang diginakan untuk belajar musik secara otodidak",
// 		Nodejs: true,
// 		Reactjs: false,
// 		JavaScript: true,
// 		Golang: true,
// 	},
// }


func main() {
	
	route := echo.New()

	connection.DatabaseConnection()

	route.Static("/assets", "assets")

	route.GET("/", home)
	route.GET("/contact", contact)
	route.GET("/project", project)
	route.GET("/testimonial", testimonial)
	route.GET("/detail-project/:id", detailProject)
	route.GET("/form-project", formProject)
	route.GET("/formUpdate/:id", formUpdate)


	route.POST("/add-project", addProject)
	route.POST("/deleteProject/:id", deleteProject)
	route.POST("/UpdateProject", updateProject)

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

	dataProjects, errDb := connection.Conn.Query(context.Background(), "select * from tb_project")

	if errDb != nil {
		return c.JSON(http.StatusInternalServerError, errDb.Error())
	}

	var resultDb []Project
	for dataProjects.Next() {
		var dataDb = Project{}

		err := dataProjects.Scan(&dataDb.Id, &dataDb.ProjectName, &dataDb.StartDate, &dataDb.EndDate, &dataDb.Duration, &dataDb.Description, &dataDb.Nodejs, &dataDb.Reactjs, &dataDb.JavaScript, &dataDb.Golang, &dataDb.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resultDb = append(resultDb, dataDb)
	}

	data := map[string]interface{} {
		"Projects" : resultDb,
	}

	return temp.Execute(c.Response(), data)
}

func testimonial(c echo.Context) error {
	temp, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}

func formProject(c echo.Context) error {
	temp, err := template.ParseFiles("views/formProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	return temp.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("nameProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := durationDistance(startDate, endDate)
	description := c.FormValue("desc")
	nodejs := c.FormValue("node") == "true"
	reactjs := c.FormValue("react") == "true"
	javaScript := c.FormValue("javaScript") == "true"
	golang := c.FormValue("go") == "true"
	image := c.FormValue("image")

	takeData, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_name, start_date, end_date, description, duration, nodejs, reactjs, java_script, golang, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", projectName, startDate, endDate, description, duration, nodejs, reactjs, javaScript, golang, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(takeData)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func detailProject(c echo.Context) error {
	id := c.Param("id")

	temp, err := template.ParseFiles("views/project-detail1.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	var detailProject = Project{}

	errDetail := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", idToInt).Scan(&detailProject.Id, &detailProject.ProjectName, &detailProject.StartDate, &detailProject.EndDate, &detailProject.Duration, &detailProject.Description, &detailProject.Nodejs, &detailProject.Reactjs, &detailProject.JavaScript, &detailProject.Golang, &detailProject.Image,)

	if errDetail != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": detailProject,
	}

	return temp.Execute(c.Response(), data)
}

func formUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/formUpdate.html")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Id": id,
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProject(c echo.Context) error {
	id := c.FormValue("id")
	projectName := c.FormValue("nameProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := durationDistance(startDate, endDate)
	description := c.FormValue("desc")
	nodejs := c.FormValue("node") == "true"
	reactjs := c.FormValue("react") == "true"
	javaScript := c.FormValue("javaScript") == "true"
	golang := c.FormValue("go") == "true"
	image := c.FormValue("image")

	sdate, _ := time.Parse("2006-01-02", startDate)
	edate, _ := time.Parse("2006-01-02", endDate)

	updated, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, duration=$4, description=$5, nodejs=$6, reactjs=$7, java_script=$8, golang=$9, image=$10 WHERE id=$11", projectName, sdate, edate, duration, description, nodejs, reactjs, javaScript, golang, image, id)

	if err != nil {
		fmt.Println("error guys", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("halo bang", updated.RowsAffected())
	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func durationDistance(dStart string, dEnd string) string {
	dateStart, _ := time.Parse("2006-01-02", dStart)
	dateEnd, _ := time.Parse("2006-01-02", dEnd)

	distance := dateEnd.Sub(dateStart)

	days := int(distance.Hours() / 24)
	weeks := days / 7
	months := weeks / 4
	years := months / 12

	if months > 12 {
		return strconv.Itoa(years) + " tahun " + strconv.Itoa(months%12) + " bulan " + strconv.Itoa(weeks%4) + " minggu " + strconv.Itoa(days%7) + " hari"
	} else if weeks > 4 {
		return strconv.Itoa(months) + " bulan " + strconv.Itoa(weeks%4) + " minggu " + strconv.Itoa(days%7) + " hari"
	} else if days >= 7 {
		return strconv.Itoa(weeks) + " minggu " + strconv.Itoa(days%7) + " hari"
	} else {
		return strconv.Itoa(days) + " hari"
	}
}

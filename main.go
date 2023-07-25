package main

import (
	"context"
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

var dataProjects = []Project{
	{
		Id: 1,
		ProjectName : "Education Mobile App",
		StartDate : time.Now(),
		EndDate : time.Now(),
		Duration : "3 minggu",
		Description : "Apalikasi ini dibuat dri tahun kmren yang diginakan untuk belajar online",
		Nodejs: true,	
		Reactjs: true,
		JavaScript: true,
		Golang: false,
	},
	{
		Id: 2,
		ProjectName : "Music Mobile App",
		StartDate : time.Now(),
		EndDate : time.Now(),
		Duration : "4 minggu 2 hari",
		Description : "Apalikasi ini dibuat dri tahun kmren yang diginakan untuk belajar musik secara otodidak",
		Nodejs: true,
		Reactjs: false,
		JavaScript: true,
		Golang: true,
	},
}


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

	dataProjects, errDb := connection.Conn.Query(context.Background(), "select * from tb_projects")

	if errDb != nil {
		return c.JSON(http.StatusInternalServerError, errDb.Error())
	}

	var resultDb []Project
	for dataProjects.Next() {
		var dataDb = Project{}

		err := dataProjects.Scan(&dataDb.ProjectName, &dataDb.StartDate, &dataDb.EndDate, &dataDb.Duration, &dataDb.Description, &dataDb.Nodejs, &dataDb.Reactjs, &dataDb.JavaScript, &dataDb.Golang, &dataDb.Image, &dataDb.Id)
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
	description := c.FormValue("desc")
	nodejs := c.FormValue("node")
	reactjs := c.FormValue("react")
	javaScript := c.FormValue("javaScript")
	golang := c.FormValue("go")
	image := c.FormValue("image")
		

	newProjectData := Project{
		ProjectName: projectName,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Duration: durationDistance(startDate, endDate),
		Description: description,
		Nodejs: (nodejs == "nodejs"),
		Reactjs: (reactjs == "reactjs"),
		JavaScript: (javaScript == "javaScript"),
		Golang: (golang == "golang"),
		Image: image,
	}


	dataProjects = append(dataProjects, newProjectData)

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

	dataProject, err := connection.Conn.Query(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", idToInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer dataProject.Close()

	if dataProject.Next() {
		var detailProject = Project{}

		err = dataProject.Scan(
			&detailProject.ProjectName,
			&detailProject.StartDate,
			&detailProject.EndDate,
			&detailProject.Duration,
			&detailProject.Description,
			&detailProject.Nodejs,
			&detailProject.Reactjs,
			&detailProject.JavaScript,
			&detailProject.Golang,
			&detailProject.Image,
			&detailProject.Id,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		data := map[string]interface{}{
			"id":      id,
			"Project": detailProject,
		}

		return temp.Execute(c.Response(), data)
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message": "Project not found"})
}


func formUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	projectUpdate := Project{}

	for index, data := range dataProjects {
		if id == index {
			projectUpdate = Project {
				Id: 			index,
				ProjectName: 	data.ProjectName,
				StartDate: 		data.StartDate,
				EndDate: 		data.EndDate,
				Duration: 		data.Duration,
				Description: 	data.Description,
				Nodejs: 		data.Nodejs,
				Reactjs:		data.Reactjs,
				JavaScript: 	data.JavaScript,
				Golang: 		data.Golang,
			}
		}
	}
	data := map[string]interface{} {
		"Project" : projectUpdate,
	}
	temp, err := template.ParseFiles("views/formUpdate.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return temp.Execute(c.Response(), data)
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	projectName := c.FormValue("nameProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("desc")
	nodejs := c.FormValue("node")
	reactjs := c.FormValue("react")
	javaScript := c.FormValue("javaScript")
	golang := c.FormValue("go")
	image := c.FormValue("image")

	updatedData := Project{
		ProjectName: projectName,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Duration: durationDistance(startDate, endDate),
		Description: description,
		Nodejs: (nodejs == "nodejs"),
		Reactjs: (reactjs == "reactjs"),
		JavaScript: (javaScript == "javaScript"),
		Golang: (golang == "golang"),
		Image: image,
	}
	dataProjects[id] = updatedData
		return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

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

	if days < 7 {
		return strconv.Itoa(days) + " hari"
	}
	if days >= 7 {
		return strconv.Itoa(weeks) + " minggu " + strconv.Itoa(days % 7) + " hari" 
	}
	if weeks >= 4 {
		return strconv.Itoa(months) + " bulan " + strconv.Itoa(weeks % 7) + " minggu " + strconv.Itoa(days % 7) + " hari"
	}
	if months >= 12 {
		return strconv.Itoa(years) + " tahun " + strconv.Itoa(months % 12) + " bulan " + strconv.Itoa(weeks % 7) + " minggu " + strconv.Itoa(days % 7) + " hari"
	} 
	return strconv.Itoa(int(distance)) + " jam"
}


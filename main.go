package main

import (
	"context"
	"fmt"
	"html/template"
	"my-web-module/connection"
	"my-web-module/middleware"
	"net/http"
	"strconv"
	"time"
	"unicode"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	Author  string
}

type User struct {
	Id int
	Name string
	Email string
	HashedPassword string
}

type UserLoginSession struct {
	IsLogin bool
	Name string
	Id int
}

var userLoginSession = UserLoginSession{}


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

	// database connetion
	connection.DatabaseConnection()

	route.Use(session.Middleware(sessions.NewCookieStore([]byte("rizky123"))))

	route.Static("/assets", "assets")
	route.Static("/uploads", "uploads")

	route.GET("/", home)
	route.GET("/contact", contact)
	route.GET("/project", project)
	route.GET("/testimonial", testimonial)
	route.GET("/detail-project/:id", detailProject)
	route.GET("/form-project", formProject)
	route.GET("/form-Update/:id", formUpdate)

	// auth
	route.GET("/form-login", formLogin)
	route.POST("/login", login)

	route.POST("/logout", logout)

	route.GET("/form-register", formRegister)
	route.POST("/register", register)


	route.POST("/add-project", middleware.UploadFile(addProject))
	route.POST("/deleteProject/:id", deleteProject)
	route.POST("/UpdateProject", middleware.UploadFile(updateProject))

	route.Logger.Fatal(route.Start("localhost:5000"))
}

func home(c echo.Context) error {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	flash := map[string]interface{} {
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}
	dataResponse := map[string]interface{}{
		"Flash": flash,
		"UserLoginSession": userLoginSession,
	}

	return temp.Execute(c.Response(), dataResponse)
}

func contact(c echo.Context) error {
	temp, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}
	dataResponse := map[string]interface{}{
		"UserLoginSession": userLoginSession,
	}

	return temp.Execute(c.Response(), dataResponse)
}

func project(c echo.Context) error {
	temp, err := template.ParseFiles("views/project.html")


	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	dataProjects, errDb := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_user.name, tb_project.project_name, tb_project.start_date, tb_project.end_date, tb_project.duration, tb_project.description, tb_project.nodejs, tb_project.reactjs, tb_project.java_script, tb_project.golang, tb_project.image FROM tb_project LEFT JOIN tb_user ON tb_project.author_id = tb_user.id")

	if errDb != nil {
		return c.JSON(http.StatusInternalServerError, errDb.Error())
	}

	var resultDb []Project
	for dataProjects.Next() {
		var dataDb = Project{}

		err := dataProjects.Scan(&dataDb.Id, &dataDb.Author, &dataDb.ProjectName, &dataDb.StartDate, &dataDb.EndDate, &dataDb.Duration, &dataDb.Description, &dataDb.Nodejs, &dataDb.Reactjs, &dataDb.JavaScript, &dataDb.Golang, &dataDb.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resultDb = append(resultDb, dataDb)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{} {
		"Projects" : resultDb,
		"UserLoginSession": userLoginSession,
	}

	return temp.Execute(c.Response(), data)
}

func testimonial(c echo.Context) error {
	temp, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}
	dataResponse := map[string]interface{}{
		"UserLoginSession": userLoginSession,
	}

	return temp.Execute(c.Response(), dataResponse)
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
	image := c.Get("dataFile").(string)

	sess, _ := session.Get("session", c)

	takeData, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_name, start_date, end_date, description, duration, nodejs, reactjs, java_script, golang, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", projectName, startDate, endDate, description, duration, nodejs, reactjs, javaScript, golang, image, sess.Values["id"].(int))

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

	errDetail := connection.Conn.QueryRow(context.Background(), "SELECT  tb_project.id, tb_user.name, tb_project.project_name, tb_project.start_date, tb_project.end_date, tb_project.duration, tb_project.description, tb_project.nodejs, tb_project.reactjs, tb_project.java_script, tb_project.golang, tb_project.image FROM tb_project LEFT JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_project.id=$1", idToInt).Scan(&detailProject.Id, &detailProject.Author, &detailProject.ProjectName, &detailProject.StartDate, &detailProject.EndDate, &detailProject.Duration, &detailProject.Description, &detailProject.Nodejs, &detailProject.Reactjs, &detailProject.JavaScript, &detailProject.Golang, &detailProject.Image)

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
	image := c.Get("dataFile").(string)

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

func formRegister(c echo.Context) error {
	temp, err := template.ParseFiles("views/registerForm.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 
	
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	flash := map[string]interface{} {
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return temp.Execute(c.Response(), flash)
}

func register(c echo.Context) error {
	
	name := c.FormValue("input-name")
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	charLenght, number, upper, special := verifyPassword(password)

	// fmt.Println("minimal banyak huruf 4:", charLenght)
    // fmt.Println("Mengandung angka:", number)
    // fmt.Println("Mengandung huruf besar:", upper)
    // fmt.Println("Mengandung karakter khusus:", special)

	if !charLenght || !number || !upper || !special {
		return redirectWithMessage(c, "Password tidak memenuhi", false, "/form-register")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataCon, errCon := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, hashedPassword)

	fmt.Println("affected rows : ", dataCon.RowsAffected())

	if errCon != nil {
		return redirectWithMessage(c, "Register gagal", false, "/form-register")
	}

	return redirectWithMessage(c, "Register Berhasil", true, "/form-login")
}

func formLogin(c echo.Context) error {
	temp, err := template.ParseFiles("views/loginForm.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} 

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	flash := map[string]interface{} {
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return temp.Execute(c.Response(), flash)
}
func login(c echo.Context) error {
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	dataUser := User{}

	errCon := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1 AND password=$2", email, password).Scan(&dataUser.Id, &dataUser.Name, &dataUser.Email, &dataUser.HashedPassword)

	if errCon != nil {
		return redirectWithMessage(c, "Email atau password salah", false, "/form-login")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(dataUser.HashedPassword), []byte(password))

	if errPass != nil {
		return redirectWithMessage(c, "Email atau password salah", false, "/form-login")
	}

	// Set login success
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return redirectWithMessage(c, "Login gagal", false, "/form-login")
	}
	sess.Options.MaxAge = 10800 
	sess.Values["message"] = "Login Berhasil"
	sess.Values["status"] = true
	sess.Values["name"] = dataUser.Name
	sess.Values["email"] = dataUser.Email
	sess.Values["id"] = dataUser.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Logout berhasil!", true, "/")
}

func verifyPassword(s string) (charLenght, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
		}
	}
	charLenght = letters >= 4
	return
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

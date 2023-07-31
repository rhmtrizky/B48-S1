package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		file, err := c.FormFile("image") //catching file byte

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("file : ", file)
		src, err := file.Open()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println("src : ", src)
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("tempFile : ", tempFile)
		defer tempFile.Close() //defer untuk menghindari dari memory list

		writtenCopy, err := io.Copy(tempFile, src)
		
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("writtenCopy : ", writtenCopy)

		data := tempFile.Name()
		fmt.Println("data : ", data)
		fileName := data[8:]

		fmt.Println("file name : ", fileName)

		c.Set("dataFile", fileName) // image-12321321.png

		return next(c)
	}
}
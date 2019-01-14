package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type User struct {
	Name  string `json:"name" xml:"name" from:"name" query:"name"`
	Email string `json:"email" xml:"email" from:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.GET("/", func(context echo.Context) error {
		fmt.Println("ssss")
		return context.String(http.StatusOK, "Hello, World!")
	})
	//文件
	e.POST("/save", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		return context.JSON(http.StatusCreated, u)
	})
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/user/:id", deleteUser)
	//e.Static("/statics", "static")
	//e.File("/statics", "static")
	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		fmt.Println("ssss")
		return context.String(http.StatusOK, "Hello, World!")
	})
	//post接收数据
	e.POST("/saveuser", saveUser)
	//文件
	e.POST("/save", save)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/user/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

//e.GET("/show", show)
//http://localhost:1323/show?team=x-men&member=wolverine
func show(context echo.Context) error {
	team := context.QueryParam("team")
	member := context.QueryParam("member")
	return context.String(http.StatusOK, "team:"+team+"，member:"+member)
}

// e.GET("/users/:id", getUser)
//http://localhost:1323/users/username
func getUser(context echo.Context) error {
	id := context.Param("id")
	fmt.Printf("%d", id)
	fmt.Println(context)
	return context.String(http.StatusOK, id)
}

//e.POST("/saveuser", saveUser)
//  curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/saveuser
func saveUser(context echo.Context) error {
	name := context.FormValue("name")
	email := context.FormValue("email")
	return context.String(http.StatusOK, "name:"+name+",email:"+email)
}

//e.POST("/save", save)
//curl -F "name=Joe Smith" -F "avatar=@C:\Users\ZhangJun\Desktop\avatar.jpg" http://localhost:1323/save
func save(context echo.Context) error {
	name := context.FormValue("name")
	avatar, err := context.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return context.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

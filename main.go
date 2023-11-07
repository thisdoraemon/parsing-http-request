package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name" form:"first_name" query:"first_name"`
	LastName  string `json:"last_name" form:"last_name" query:"last_name"`
	Email     string `json:"email" form:"email" query:"email"`
}

func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	r := echo.New()

	r.POST("/users", CreateUser)

	fmt.Println("Server started at :9000")
	err := r.Start(":9000")
	if err != nil {
		panic(err)
	}
}

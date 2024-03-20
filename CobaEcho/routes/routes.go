package routes

import (
	"CobaEcho/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is Irvan")
	})

	e.GET("/users", controllers.FetchAllUsers)
	e.POST("/users", controllers.StoreUsers)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	return e
}

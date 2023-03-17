package routes

import (
	ProductController "Product/app/Http/Controllers"
	"os"

	"github.com/labstack/echo/v4"
)

func Api() {

	e := echo.New()

	e.GET("products", ProductController.Index)
	e.GET("products/:id", ProductController.Show)
	e.POST("products", ProductController.Store)
	e.PUT("products/:id", ProductController.Update)
	e.DELETE("products/:id", ProductController.Delete)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

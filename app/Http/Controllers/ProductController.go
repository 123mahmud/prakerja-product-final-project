package ProductController

import (
	"Product/app/Models"
	DB "Product/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var products []Models.Products

func Index(c echo.Context) error {
	results := DB.Init().Find(&products)
	fmt.Println(results.Error)
	if err := results.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func Show(c echo.Context) error {

	id := c.Param("id")
	product := DB.Init().Find(&Models.Products{}, id)
	if err := product.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func Store(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	description := c.FormValue("description")
	price_int, _ := strconv.Atoi(c.FormValue("price"))

	if len(name) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "name is required",
		})
	}
	if len(price) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price is required",
		})
	}
	if len(description) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "description is required",
		})
	}
	if price_int == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price is numeric",
		})
	}
	if price_int < 1000 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price min 1000",
		})
	}

	product := Models.Products{
		Name:        name,
		Price:       price_int,
		Description: description,
	}

	result := DB.Init().Create(&product)

	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

func Update(c echo.Context) error {

	id := c.Param("id")

	name := c.FormValue("name")
	price := c.FormValue("price")
	description := c.FormValue("description")
	price_int, _ := strconv.Atoi(price)

	if len(name) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "name is required",
		})
	}
	if len(price) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price is required",
		})
	}
	if len(description) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "description is required",
		})
	}
	if price_int == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price is numeric",
		})
	}
	if price_int < 1000 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "price min 1000",
		})
	}

	product := DB.Init().Find(&Models.Products{}, id)
	if err := product.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	product.Update(Models.Products{
		Name:        name,
		Price:       price_int,
		Description: description,
	})

	return c.JSON(http.StatusOK, product)
}

func Delete(c echo.Context) error {

	id := c.Param("id")
	product := DB.Init().Find(&Models.Products{}, id)
	if err := product.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	DB.Init().Delete(&Models.Products{}, id)

	return c.JSON(http.StatusOK, product)
}

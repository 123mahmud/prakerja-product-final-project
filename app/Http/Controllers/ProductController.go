package ProductController

import (
	"Product/app/Models"
	DB "Product/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var products []Models.Products
	results := DB.Init().Find(&products)
	if err := results.Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    products,
		"message": "success products get ",
	})

	// return c.JSON(http.StatusOK, products)
}

func Show(c echo.Context) error {
	id := c.Param("id")

	product := &Models.Products{}
	result := DB.Init().Find(product, id)
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    product,
		"message": "success products show ",
	})

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    product,
		"message": "success products store",
	})

	// return c.JSON(http.StatusCreated, product)
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

	product := &Models.Products{}
	results := DB.Init().First(product, id)
	if results.RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	} else if results.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, results.Error.Error())
	}

	product.Name = name
	product.Price = price_int
	product.Description = description

	results = DB.Init().Save(product)
	if results.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, results.Error.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    product,
		"message": "success product update",
	})
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	product := &Models.Products{}
	result := DB.Init().Find(product, id)
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	err := DB.Init().Delete(product, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    product,
		"message": "success products destroy",
	})
}

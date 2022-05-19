package handlers

import (
	"final/cmd/repository"
	"final/cmd/todos"
	"final/cmd/weather"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

type API struct {
	StorageService repository.Repository
	WeatherApp     weather.WeatherInfo
}

// GetLists handler
func (api API) GetLists() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, api.StorageService.GetLists())
	}
}

// PutList handler
func (api API) PutList() echo.HandlerFunc {
	return func(c echo.Context) error {
		var List todos.List
		c.Bind(&List)
		id, err := api.StorageService.PutList(List.Name)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

// DeleteList handler
func (api API) DeleteList() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := api.StorageService.DeleteList(id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (api API) GetWeather() echo.HandlerFunc {
	return func(c echo.Context) error {
		//get parameters from api
		latitude := c.Request().Header.Get("lat")
		longitude := c.Request().Header.Get("lon")
		fmt.Println(latitude)
		fmt.Println(longitude)
		return c.JSON(http.StatusOK, api.WeatherApp.GetWeather(latitude, longitude))
	}
}

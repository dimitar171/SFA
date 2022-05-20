package handlers

import (
	"final/cmd/weather"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WeatherAPI struct {
	WeatherUrl weather.WeatherUrl
}

func (apiW WeatherAPI) GetWeather() echo.HandlerFunc {
	return func(c echo.Context) error {
		//get parameters from api
		latitude := c.Request().Header.Get("lat")
		longitude := c.Request().Header.Get("lon")
		return c.JSON(http.StatusOK, apiW.WeatherUrl.GetWeather(latitude, longitude))
	}
}

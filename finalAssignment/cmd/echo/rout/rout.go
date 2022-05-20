package rout

import (
	"final/cmd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(apiH handlers.API, apiW handlers.WeatherAPI) *echo.Echo {
	router := echo.New()

	router.Use(middleware.BasicAuth(apiH.Authorize))

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})
	router.GET("/api/list/export", apiH.GetCsv())
	router.GET("/api/weather", apiW.GetWeather())
	router.DELETE("/api/lists/:id", apiH.DeleteList())
	router.POST("/api/lists", apiH.PutList())
	router.GET("/api/lists", apiH.GetLists())
	router.DELETE("/api/tasks/:id", apiH.DeleteTask())
	router.POST("/api/lists/:id/tasks", apiH.PutTask())
	router.GET("/api/lists/:id/tasks", apiH.GetTasks())
	router.PATCH("/api/tasks/:id", apiH.PatchTasks())

	return router
}

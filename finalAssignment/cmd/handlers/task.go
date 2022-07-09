package handlers

import (
	"final/cmd/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetTasks endpoint
func (api API) GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, api.StorageService.GetTasks(id))
	}
}

// PutTask endpoint
func (api API) PutTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Task repository.Task
		c.Bind(&Task)
		id1, _ := strconv.Atoi(c.Param("id"))
		id, err := api.StorageService.PutTask(Task.Text, id1, Task.Completed)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func (api API) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := api.StorageService.DeleteTask(id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}

// PatchTasks endpoint
func (api API) PatchTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := api.StorageService.PatchTask(id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": id,
			})
		} else {
			return err
		}
	}
}

package handlers

import (
	"bytes"
	"encoding/csv"
	"final/cmd/repository"
	"final/cmd/todos"
	"final/cmd/weather"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

type API struct {
	StorageService repository.Repository
	WeatherApp     weather.WeatherInfo
}

func (api API) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetLists endpoint
func (api API) GetLists() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, api.StorageService.GetLists())
	}
}

// PutList endpoint
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

// DeleteList endpoint
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
		var Task todos.Task
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

func (api API) Authorize(username, password string, c echo.Context) (bool, error) {
	userStruct := api.StorageService.GetUsers()
	for _, v := range userStruct {
		hashPassword := api.CheckPasswordHash(password, v.Pass)
		if username == v.Name && hashPassword {
			api.StorageService.PatchCurrentUser(v.ID)
			return true, nil
		}
	}
	return false, nil
}

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

type Contact struct {
	Email string
	Open  int64
	Link  int64
}

func (api API) GetCsv() echo.HandlerFunc {
	return func(c echo.Context) error {
		Lists := api.StorageService.GetLists()
		var data [][]string
		for v, list := range Lists {
			row := []string{strconv.Itoa(list.ID), list.Name, strconv.Itoa(list.UserId)}
			data = append(data, row)
			Tasks := api.StorageService.GetTasks(v)
			for _, tasks := range Tasks {
				row := []string{strconv.Itoa(tasks.Id), tasks.Text, strconv.Itoa(tasks.ListId), strconv.FormatBool(tasks.Completed)}
				data = append(data, row)
			}
		}
		b := &bytes.Buffer{}
		wr := csv.NewWriter(b)
		wr.WriteAll(data)
		wr.Flush()
		c.Response().Writer.Header().Set("Content-Description", "File Transfer")
		c.Response().Writer.Header().Set("Content-Disposition", "attachment; filename=list.csv")

		return c.JSON(c.Response().Writer.Write(b.Bytes()))
	}
}

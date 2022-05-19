package handlers

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

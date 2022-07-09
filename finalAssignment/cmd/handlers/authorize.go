package handlers

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func (api API) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

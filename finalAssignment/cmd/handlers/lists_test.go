package handlers

// import (
// 	"database/sql"
// 	"final/cmd/repository"
// 	"final/cmd/todos"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo"
// 	"github.com/stretchr/testify/assert"
// )

// var (

// 	mockDB   = []todos.List{{1, "List1"}}
// 	listJSON = `{1, "List1"}`
// )

// func TestPutList(t *testing.T) {
// 	mockDb, _ := sql.Open("sqlite", ":memory:")
// 	mockRepo:=repository.NewRepository(mockDb)
// 	MockApi:=API{StorageService: *mockRepo}
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(listJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	// Assertions
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 		assert.Equal(t, listJSON, rec.Body.String())

// }

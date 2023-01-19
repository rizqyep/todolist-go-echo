package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var todoJSON = `{"task":"create test"}`

func TestCreateTodo(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := NewHandlersInstance().Todo
	var resultJSON map[string]map[string]string
	json.Unmarshal(rec.Body.Bytes(), &resultJSON)
	var expectedJSON map[string]string
	json.Unmarshal([]byte(todoJSON), &expectedJSON)
	if assert.NoError(t, h.CreateTodo(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

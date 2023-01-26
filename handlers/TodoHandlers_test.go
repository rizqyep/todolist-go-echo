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

var todoJSON = `{"task":"testing task"}`

type TodoResponse struct {
	ResponseStructure
	Data json.RawMessage
}

func TestCreateTodo(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := InitHandlerInstance().Todo

	assert.NoError(t, h.CreateTodo(c))

}

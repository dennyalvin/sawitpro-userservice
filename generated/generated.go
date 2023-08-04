package generated

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func WrapResponseJson(ctx echo.Context, title string, data interface{}) error {
	resp := APIResponse{
		Message: title,
		Data:    data,
	}
	return ctx.JSON(http.StatusOK, resp)
}

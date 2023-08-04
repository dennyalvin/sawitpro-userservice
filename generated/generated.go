package generated

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func WrapResponseJsonOK(ctx echo.Context, title string, data interface{}) error {
	resp := APIResponse{
		Message: title,
		Data:    data,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func WrapResponseJsonBadRequest(ctx echo.Context, errDtails []ErrorDetail) error {
	resp := ResponseError{
		Message: "Bad Request",
		Errors:  errDtails,
	}
	return ctx.JSON(http.StatusBadRequest, resp)
}

package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) GetProfileDetail(c echo.Context, id int) error {
	var resp interface{}

	//call repository to find by id
	user, err := s.Repository.FindBy(c.Request().Context(), "id", id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "please contact administrator")
	}

	//if id is not found in db, return 403 forbidden status
	if user == nil {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	//create object for restful response if user is founded
	resp = generated.ProfileResponse{
		FullName: user.FullName,
		Phone:    user.Phone,
	}

	return generated.WrapResponseJson(c, "OK", resp)
}

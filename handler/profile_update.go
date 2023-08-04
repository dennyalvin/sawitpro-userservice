package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) UpdateProfile(c echo.Context, userId int, params *generated.ProfileUpdateParams) error {
	var resp interface{}

	//call repo update
	err := s.Repository.Update(c.Request().Context(), userId, params)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "please contact administrator")
	}

	return generated.WrapResponseJson(c, "profile updated", resp)
}

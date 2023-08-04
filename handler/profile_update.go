package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) UpdateProfile(c echo.Context, userId int, params generated.ProfileUpdateParams) error {
	var resp interface{}

	//Validate Request
	errValidation := s.ValidateUpdateProfileRequest(params)
	if errValidation != nil {
		return generated.WrapResponseJsonBadRequest(c, errValidation)
	}

	//call repo to update
	err := s.Repository.Update(c.Request().Context(), userId, params)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "please contact administrator")
	}

	return generated.WrapResponseJsonOK(c, "profile updated", resp)
}

func (s *Server) ValidateUpdateProfileRequest(params generated.ProfileUpdateParams) []generated.ErrorDetail {
	var errs []generated.ErrorDetail

	errs = ValidateStruct(params)
	if errs != nil {
		return errs
	}

	user, _ := s.Repository.FindBy(context.Background(), "phone", params.Phone)
	if user != nil {
		errs = append(errs, generated.ErrorDetail{
			Title:   "phone",
			Message: "is already exists",
		})

		return errs
	}

	return errs
}

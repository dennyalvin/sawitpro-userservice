package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Server) Register(c echo.Context, params generated.SignupParams) error {

	//Validate Request
	errValidation := s.ValidateRegistrationRequest(params)
	if errValidation != nil {
		return generated.WrapResponseJsonBadRequest(c, errValidation)
	}

	//Hash Password
	hashedPassword, err := hashPassword(params.Password)
	if err != nil {
		return echo.ErrInternalServerError
	}

	//prepare object to insert into db
	user := model.User{
		FullName: params.FullName,
		Password: hashedPassword,
		Phone:    params.Phone,
	}

	//call create repository
	newUser, err := s.Repository.Create(c.Request().Context(), user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "please contact administrator")
	}

	//create object response
	resp := generated.ProfileResponse{
		FullName: newUser.FullName,
		Phone:    newUser.Phone,
	}

	return generated.WrapResponseJsonOK(c, "New user successfully created", resp)
}

func hashPassword(password string) (string, error) {
	saltedPass := helper.SaltString(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *Server) ValidateRegistrationRequest(params generated.SignupParams) []generated.ErrorDetail {
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

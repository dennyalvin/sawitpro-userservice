package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Server) Register(c echo.Context, params *generated.SignupParams) error {

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

	return generated.WrapResponseJson(c, "New user successfully created", resp)
}

func hashPassword(password string) (string, error) {
	saltedPass := helper.SaltString(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

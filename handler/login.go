package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Server) Login(c echo.Context, params *generated.LoginParams) error {
	ctx := c.Request().Context()
	//Find user by Phone
	user, err := s.Repository.FindBy(ctx, "phone", params.Phone)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "please contact administrator")
	}

	//If phone number is not founded
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "phone number or password is incorrect")
	}

	//If user is found, then compare the password
	if !isPasswordValid(user.Password, params.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "phone number or password is incorrect")
	}

	//Then increment the success login counter
	err = s.Repository.UpdateLoginSuccess(ctx, user.Id)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}

	//Generate the JWT Signed token
	token, err := helper.GetJWTToken(user.Id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	//Generate Restful Response for success login
	auth := generated.LoginResponse{
		Id:    user.Id,
		Token: token,
	}

	return generated.WrapResponseJson(c, "Login successful", auth)
}

func isPasswordValid(hashedPassword string, plainPassword string) bool {
	saltedPass := helper.SaltString(plainPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(saltedPass))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

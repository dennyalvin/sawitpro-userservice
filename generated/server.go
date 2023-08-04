package generated

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServerInterface interface {
	Register(c echo.Context, params *SignupParams) error
	Login(c echo.Context, params *LoginParams) error
	UpdateProfile(c echo.Context, id int, params *ProfileUpdateParams) error
	GetProfileDetail(c echo.Context, id int) error
}

func RegisterHandlers(e *echo.Echo, s ServerInterface) {

	//Group endpoint path /api/users
	api := e.Group("/api")
	user := api.Group("/users")

	//Set Router and bind request params
	user.POST("/register", func(c echo.Context) error {
		var params SignupParams

		if err := generateParams(c, &params); err != nil {
			return err
		}

		return s.Register(c, &params)
	})

	user.POST("/login", func(c echo.Context) error {
		var params LoginParams

		if err := generateParams(c, &params); err != nil {
			return err
		}

		return s.Login(c, &params)
	})

	// route with auth jwt middleware
	user.GET("/show", func(c echo.Context) error {
		token := c.Get("user")
		userId := helper.ClaimJWTUserId(token)
		return s.GetProfileDetail(c, userId)
	}, middlewareEchoJWT())

	user.PATCH("", func(c echo.Context) error {
		var params ProfileUpdateParams
		token := c.Get("user")
		userId := helper.ClaimJWTUserId(token)

		if err := generateParams(c, &params); err != nil {
			return err
		}

		return s.UpdateProfile(c, userId, &params)
	}, middlewareEchoJWT())
}

func generateParams(c echo.Context, structParams interface{}) error {
	if err := c.Bind(structParams); err != nil {
		return err
	}

	return nil
}

func middlewareEchoJWT() echo.MiddlewareFunc {

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: helper.JWTSecretKey,
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Print("sdfd")
			// as assignment expected need to return Forbidden status code
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
		},
	})
}

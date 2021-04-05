package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/helpers"
)

func CheckAuthorization(c *fiber.Ctx) error {
	clientToken := c.Get("Authorization")
	if clientToken == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnauthorized,
			Message: "NO-AUTHORIZATION-HEADER-PROVIDED",
			Data:    nil,
		})
	}
	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		return c.JSON(config.AppResponse{
			Code:    400,
			Message: "INCORRECT-FORMAT-AUTHORIZATION",
			Data:    nil,
		})
	}

	jwtWrapper := helpers.JwtWrapper{
		SecretKey: config.GoDotEnvVariable("SECRET_KEY"),
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(clientToken)
	if err != nil {
		return c.JSON(config.AppResponse{
			Code:    401,
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.Set("id", claims.ID)
	c.Set("email", claims.Email)

	return c.Next()
}

package apiController

import (
	"github.com/labstack/echo/v4"
	apiHelper "github.com/speedianet/os/src/presentation/api/helper"
	"github.com/speedianet/os/src/presentation/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}

// GenerateJwtWithCredentials godoc
// @Summary      GenerateJwtWithCredentials
// @Description  Generate JWT with credentials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginDto 	  body    dto.Login  true  "All props are required."
// @Success      200 {object} entity.AccessToken
// @Failure      401 {object} string
// @Router       /v1/auth/login/ [post]
func (controller *AuthController) GenerateJwtWithCredentials(c echo.Context) error {
	requestBody, err := apiHelper.ReadRequestBody(c)
	if err != nil {
		return err
	}
	requestBody["ipAddress"] = c.RealIP()

	return apiHelper.ServiceResponseWrapper(
		c, controller.authService.GenerateJwtWithCredentials(requestBody),
	)
}

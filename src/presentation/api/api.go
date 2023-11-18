package api

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apiInit "github.com/speedianet/os/src/presentation/api/init"
	apiMiddleware "github.com/speedianet/os/src/presentation/api/middleware"
	sharedMiddleware "github.com/speedianet/os/src/presentation/shared/middleware"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title			SosApi
// @version			0.0.1
// @description		Speedia OS API
// @termsOfService	https://speedia.net/tos/

// @contact.name	Speedia Engineering
// @contact.url		https://speedia.net/
// @contact.email	eng+swagger@speedia.net

// @license.name  Eclipse Public License v2.0
// @license.url   https://www.eclipse.org/legal/epl-2.0/

// @securityDefinitions.apikey	Bearer
// @in 							header
// @name						Authorization
// @description					Type "Bearer" + JWT token or API key.

// @host		localhost:10000
// @BasePath	/v1
func ApiInit() {
	sharedMiddleware.CheckEnvs()
	apiInit.WebServerSetup()

	e := echo.New()

	basePath := "/v1"
	baseRoute := e.Group(basePath)

	requestTimeout := 60 * time.Second

	e.Pre(apiMiddleware.TrailingSlash(basePath))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: requestTimeout,
	}))
	e.Use(apiMiddleware.PanicHandler)
	e.Use(apiMiddleware.SetDefaultHeaders)
	e.Use(apiMiddleware.Auth(basePath))

	registerApiRoutes(baseRoute)

	e.Start(":10000")
}

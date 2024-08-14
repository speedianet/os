package ui

import (
	"github.com/labstack/echo/v4"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	uiMiddleware "github.com/speedianet/os/src/presentation/ui/middleware"
)

func UiInit(
	e *echo.Echo,
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) {
	basePath := ""
	baseRoute := e.Group(basePath)

	e.Use(uiMiddleware.Authentication())

	router := NewRouter(baseRoute, persistentDbSvc)
	router.RegisterRoutes()
}

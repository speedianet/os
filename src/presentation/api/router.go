package api

import (
	_ "embed"

	"github.com/labstack/echo/v4"
	apiController "github.com/speedianet/os/src/presentation/api/controller"
	apiMiddleware "github.com/speedianet/os/src/presentation/api/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/speedianet/os/src/presentation/api/docs"
)

func swaggerRoute(baseRoute *echo.Group) {
	swaggerGroup := baseRoute.Group("/swagger")
	swaggerGroup.GET("/*", echoSwagger.WrapHandler)
}

func authRoutes(baseRoute *echo.Group) {
	authGroup := baseRoute.Group("/auth")
	authGroup.POST("/login/", apiController.AuthLoginController)
}

func accountRoutes(baseRoute *echo.Group) {
	accountGroup := baseRoute.Group("/account")
	accountGroup.GET("/", apiController.GetAccountsController)
	accountGroup.POST("/", apiController.AddAccountController)
	accountGroup.PUT("/", apiController.UpdateAccountController)
	accountGroup.DELETE("/:accountId/", apiController.DeleteAccountController)
}

func cronRoutes(baseRoute *echo.Group) {
	cronGroup := baseRoute.Group("/cron")
	cronGroup.GET("/", apiController.GetCronsController)
	cronGroup.POST("/", apiController.AddCronController)
	cronGroup.PUT("/", apiController.UpdateCronController)
	cronGroup.DELETE("/:cronId/", apiController.DeleteCronController)
}

func databaseRoutes(baseRoute *echo.Group) {
	databaseGroup := baseRoute.Group("/database", apiMiddleware.ServiceStatusValidator("mysql"))
	databaseGroup.GET("/:dbType/", apiController.GetDatabasesController)
	databaseGroup.POST("/:dbType/", apiController.AddDatabaseController)
	databaseGroup.DELETE(
		"/:dbType/:dbName/",
		apiController.DeleteDatabaseController,
	)
	databaseGroup.POST(
		"/:dbType/:dbName/user/",
		apiController.AddDatabaseUserController,
	)
	databaseGroup.DELETE(
		"/:dbType/:dbName/user/:dbUser/",
		apiController.DeleteDatabaseUserController,
	)
}

func filesRoutes(baseRoute *echo.Group) {
	filesGroup := baseRoute.Group("/files")
	filesGroup.GET("/", apiController.GetFilesController)
	filesGroup.POST("/", apiController.AddFileController)
	filesGroup.PUT("/", apiController.UpdateFileController)
	filesGroup.POST("/copy/", apiController.AddFileCopyController)
	filesGroup.PUT("/content/", apiController.UpdateFileContentController)
	filesGroup.PUT("/delete/", apiController.DeleteFileController)
}

func o11yRoutes(baseRoute *echo.Group) {
	o11yGroup := baseRoute.Group("/o11y")
	o11yGroup.GET("/overview/", apiController.O11yOverviewController)
}

func runtimeRoutes(baseRoute *echo.Group) {
	runtimeGroup := baseRoute.Group("/runtime")
	runtimeGroup.GET("/php/:hostname/", apiController.GetPhpConfigsController)
	runtimeGroup.PUT("/php/:hostname/", apiController.UpdatePhpConfigsController)
}

func servicesRoutes(baseRoute *echo.Group) {
	servicesGroup := baseRoute.Group("/services")
	servicesGroup.GET("/", apiController.GetServicesController)
	servicesGroup.PUT("/", apiController.UpdateServiceController)
}

func sslRoutes(baseRoute *echo.Group) {
	sslGroup := baseRoute.Group("/ssl")
	sslGroup.GET("/", apiController.GetSslPairsController)
	sslGroup.POST("/", apiController.AddSslPairController)
	sslGroup.DELETE("/:sslPairId/", apiController.DeleteSslPairController)
}

func registerApiRoutes(baseRoute *echo.Group) {
	swaggerRoute(baseRoute)
	authRoutes(baseRoute)
	accountRoutes(baseRoute)
	cronRoutes(baseRoute)
	databaseRoutes(baseRoute)
	filesRoutes(baseRoute)
	o11yRoutes(baseRoute)
	runtimeRoutes(baseRoute)
	servicesRoutes(baseRoute)
	sslRoutes(baseRoute)
}

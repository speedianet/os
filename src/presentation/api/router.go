package api

import (
	_ "embed"

	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	apiController "github.com/goinfinite/os/src/presentation/api/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/goinfinite/os/src/presentation/api/docs"
)

type Router struct {
	baseRoute       *echo.Group
	transientDbSvc  *internalDbInfra.TransientDatabaseService
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
	trailDbSvc      *internalDbInfra.TrailDatabaseService
}

func NewRouter(
	baseRoute *echo.Group,
	transientDbSvc *internalDbInfra.TransientDatabaseService,
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
	trailDbSvc *internalDbInfra.TrailDatabaseService,
) *Router {
	return &Router{
		baseRoute:       baseRoute,
		transientDbSvc:  transientDbSvc,
		persistentDbSvc: persistentDbSvc,
		trailDbSvc:      trailDbSvc,
	}
}

func (router Router) swaggerRoute() {
	swaggerGroup := router.baseRoute.Group("/swagger")
	swaggerGroup.GET("/*", echoSwagger.WrapHandler)
}

func (router Router) authRoutes() {
	authGroup := router.baseRoute.Group("/v1/auth")
	authController := apiController.NewAuthController(
		router.persistentDbSvc, router.trailDbSvc,
	)

	authGroup.POST("/login/", authController.Login)
}

func (router Router) accountRoutes() {
	accountGroup := router.baseRoute.Group("/v1/account")
	accountController := apiController.NewAccountController(
		router.persistentDbSvc, router.trailDbSvc,
	)

	accountGroup.GET("/", accountController.Read)
	accountGroup.POST("/", accountController.Create)
	accountGroup.PUT("/", accountController.Update)
	accountGroup.DELETE("/:accountId/", accountController.Delete)
}

func (router Router) cronRoutes() {
	cronGroup := router.baseRoute.Group("/v1/cron")
	cronController := apiController.NewCronController()

	cronGroup.GET("/", cronController.Read)
	cronGroup.POST("/", cronController.Create)
	cronGroup.PUT("/", cronController.Update)
	cronGroup.DELETE("/:cronId/", cronController.Delete)
}

func (router Router) databaseRoutes() {
	databaseGroup := router.baseRoute.Group("/v1/database")
	databaseController := apiController.NewDatabaseController(
		router.persistentDbSvc,
	)

	databaseGroup.GET("/:dbType/", databaseController.Read)
	databaseGroup.POST("/:dbType/", databaseController.Create)
	databaseGroup.DELETE("/:dbType/:dbName/", databaseController.Delete)
	databaseGroup.POST("/:dbType/:dbName/user/", databaseController.CreateUser)
	databaseGroup.POST("/:dbType/user/", databaseController.CreateUser)
	databaseGroup.DELETE(
		"/:dbType/:dbName/user/:dbUser/", databaseController.DeleteUser,
	)
}

func (router Router) filesRoutes() {
	filesGroup := router.baseRoute.Group("/v1/files")
	filesGroup.GET("/", apiController.GetFilesController)
	filesGroup.POST("/", apiController.CreateFileController)
	filesGroup.PUT("/", apiController.UpdateFileController)
	filesGroup.POST("/copy/", apiController.CopyFileController)
	filesGroup.PUT("/delete/", apiController.DeleteFileController)
	filesGroup.POST("/compress/", apiController.CompressFilesController)
	filesGroup.PUT("/extract/", apiController.ExtractFilesController)
	filesGroup.POST("/upload/", apiController.UploadFilesController)
}

func (router Router) marketplaceRoutes() {
	marketplaceGroup := router.baseRoute.Group("/v1/marketplace")
	marketplaceController := apiController.NewMarketplaceController(
		router.persistentDbSvc,
	)

	marketplaceInstalledGroup := marketplaceGroup.Group("/installed")
	marketplaceInstalledGroup.GET("/", marketplaceController.ReadInstalledItems)
	marketplaceInstalledGroup.DELETE(
		"/:installedId/",
		marketplaceController.DeleteInstalledItem,
	)

	marketplaceCatalogGroup := marketplaceGroup.Group("/catalog")
	marketplaceCatalogGroup.GET("/", marketplaceController.ReadCatalog)
	marketplaceCatalogGroup.POST("/", marketplaceController.InstallCatalogItem)
}

func (router Router) o11yRoutes() {
	o11yGroup := router.baseRoute.Group("/v1/o11y")

	o11yController := apiController.NewO11yController(router.transientDbSvc)
	o11yGroup.GET("/overview/", o11yController.ReadOverview)
}

func (router Router) runtimeRoutes() {
	runtimeGroup := router.baseRoute.Group("/v1/runtime")
	runtimeController := apiController.NewRuntimeController(
		router.persistentDbSvc,
	)

	runtimeGroup.GET("/php/:hostname/", runtimeController.ReadPhpConfigs)
	runtimeGroup.PUT("/php/:hostname/", runtimeController.UpdatePhpConfigs)
}

func (router *Router) scheduledTaskRoutes() {
	scheduledTaskGroup := router.baseRoute.Group("/v1/scheduled-task")

	scheduledTaskController := apiController.NewScheduledTaskController(router.persistentDbSvc)
	scheduledTaskGroup.GET("/", scheduledTaskController.Read)
	scheduledTaskGroup.PUT("/", scheduledTaskController.Update)
	go scheduledTaskController.Run()
}

func (router Router) servicesRoutes() {
	servicesGroup := router.baseRoute.Group("/v1/services")
	servicesController := apiController.NewServicesController(
		router.persistentDbSvc,
	)

	servicesGroup.GET("/", servicesController.Read)
	servicesGroup.GET("/installables/", servicesController.ReadInstallables)
	servicesGroup.POST("/installables/", servicesController.CreateInstallable)
	servicesGroup.POST("/custom/", servicesController.CreateCustom)
	servicesGroup.PUT("/", servicesController.Update)
	servicesGroup.DELETE("/:svcName/", servicesController.Delete)
}

func (router Router) sslRoutes() {
	sslGroup := router.baseRoute.Group("/v1/ssl")
	sslController := apiController.NewSslController(
		router.persistentDbSvc, router.transientDbSvc,
	)

	sslGroup.GET("/", sslController.Read)
	sslGroup.POST("/", sslController.Create)
	sslGroup.DELETE("/:sslPairId/", sslController.Delete)
	sslGroup.PUT("/vhost/", sslController.DeleteVhosts)
	go sslController.SslCertificateWatchdog()
}

func (router Router) vhostsRoutes() {
	vhostsGroup := router.baseRoute.Group("/v1/vhosts")
	vhostController := apiController.NewVirtualHostController(
		router.persistentDbSvc,
	)

	vhostsGroup.GET("/", vhostController.Read)
	vhostsGroup.POST("/", vhostController.Create)
	vhostsGroup.DELETE("/:hostname/", vhostController.Delete)

	mappingsGroup := vhostsGroup.Group("/mapping")
	mappingsGroup.GET("/", vhostController.ReadWithMappings)
	mappingsGroup.POST("/", vhostController.CreateMapping)
	mappingsGroup.DELETE(
		"/:mappingId/",
		vhostController.DeleteMapping,
	)
}

func (router Router) RegisterRoutes() {
	router.swaggerRoute()
	router.authRoutes()
	router.accountRoutes()
	router.cronRoutes()
	router.databaseRoutes()
	router.filesRoutes()
	router.marketplaceRoutes()
	router.o11yRoutes()
	router.runtimeRoutes()
	router.scheduledTaskRoutes()
	router.servicesRoutes()
	router.sslRoutes()
	router.vhostsRoutes()
}

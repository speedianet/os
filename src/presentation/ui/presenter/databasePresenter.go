package presenter

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	"github.com/speedianet/os/src/presentation/service"
	uiHelper "github.com/speedianet/os/src/presentation/ui/helper"
	"github.com/speedianet/os/src/presentation/ui/page"
	presenterDto "github.com/speedianet/os/src/presentation/ui/presenter/dto"
)

type DatabasesPresenter struct {
	databaseService *service.DatabaseService
}

func NewDatabasesPresenter(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *DatabasesPresenter {
	return &DatabasesPresenter{
		databaseService: service.NewDatabaseService(persistentDbSvc),
	}
}

func (presenter *DatabasesPresenter) getDatabaseTypeDetails(
	rawDatabaseType string,
) (databaseTypeDetails presenterDto.DatabaseTypeDetails, err error) {
	databaseType, err := valueObject.NewDatabaseType(rawDatabaseType)
	if err != nil {
		return databaseTypeDetails, err
	}

	isInstalled := false

	requestInput := map[string]interface{}{"dbType": databaseType.String()}
	responseOutput := presenter.databaseService.Read(requestInput)
	if responseOutput.Status == service.Success {
		isInstalled = true
	}

	databases, assertOk := responseOutput.Body.([]entity.Database)
	if !assertOk {
		isInstalled = false
	}

	return presenterDto.NewDatabaseTypeDetails(
		databaseType, isInstalled, databases,
	), nil
}

func (presenter *DatabasesPresenter) Handler(c echo.Context) error {
	rawDatabaseType := "mariadb"
	if c.QueryParam("selectedDatabaseType") != "" {
		rawDatabaseType = c.QueryParam("selectedDatabaseType")
	}

	databaseTypeDetails, err := presenter.getDatabaseTypeDetails(rawDatabaseType)
	if err != nil {
		slog.Debug("GetDatabaseTypeDetailsError", slog.Any("err", err))
		return nil
	}
	log.Printf("DatabaseTypeDetails: %+v", databaseTypeDetails)

	pageContent := page.DatabasesIndex(databaseTypeDetails)
	return uiHelper.Render(c, pageContent, http.StatusOK)
}

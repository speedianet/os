package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/sam/src/domain/dto"
	"github.com/speedianet/sam/src/domain/useCase"
	"github.com/speedianet/sam/src/domain/valueObject"
	"github.com/speedianet/sam/src/infra"
	apiHelper "github.com/speedianet/sam/src/presentation/api/helper"
)

// GetCrons	 godoc
// @Summary      GetCrons
// @Description  List Crons.
// @Tags         cron
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200 {array} entity.Cron
// @Router       /cron/ [get]
func GetCronsController(c echo.Context) error {
	cronsQueryRepo := infra.CronQueryRepo{}
	cronsList, err := useCase.GetCrons(cronsQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, cronsList)
}

// AddCron    godoc
// @Summary      AddNewCron
// @Description  Add a new cron.
// @Tags         cron
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        addCronDto 	  body    dto.AddCron  true  "NewCron"
// @Success      201 {object} object{} "CronCreated"
// @Router       /cron/ [post]
func AddCronController(c echo.Context) error {
	requiredParams := []string{"schedule", "command"}
	requestBody, _ := apiHelper.GetRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	var cronCommentPtr *valueObject.CronComment
	if requestBody["comment"] != nil {
		cronComment := valueObject.NewCronCommentPanic(requestBody["comment"].(string))
		cronCommentPtr = &cronComment
	}

	addCronDto := dto.NewAddCron(
		valueObject.NewCronSchedulePanic(requestBody["schedule"].(string)),
		valueObject.NewUnixCommandPanic(requestBody["command"].(string)),
		cronCommentPtr,
	)

	cronCmdRepo := infra.CronCmdRepo{}

	err := useCase.AddCron(
		cronCmdRepo,
		addCronDto,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusCreated, "CronCreated")
}

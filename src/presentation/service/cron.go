package service

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	cronInfra "github.com/speedianet/os/src/infra/cron"
	serviceHelper "github.com/speedianet/os/src/presentation/service/helper"
)

type CronService struct {
	queryRepo cronInfra.CronQueryRepo
}

func NewCronService() *CronService {
	return &CronService{
		queryRepo: cronInfra.CronQueryRepo{},
	}
}

func (service *CronService) Read() ServiceOutput {
	cronsList, err := useCase.ReadCrons(service.queryRepo)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, cronsList)
}

func (service *CronService) Create(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"schedule", "command"}
	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	schedule, err := valueObject.NewCronSchedule(input["schedule"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	command, err := valueObject.NewUnixCommand(input["command"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var commentPtr *valueObject.CronComment
	if input["comment"] != nil {
		comment, err := valueObject.NewCronComment(input["comment"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		commentPtr = &comment
	}

	dto := dto.NewCreateCron(schedule, command, commentPtr)

	cmdRepo, err := cronInfra.NewCronCmdRepo()
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	err = useCase.CreateCron(cmdRepo, dto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "CronCreated")
}

func (service *CronService) Update(input map[string]interface{}) ServiceOutput {
	requiredParams := []string{"id"}
	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	id, err := valueObject.NewCronId(input["id"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	var schedulePtr *valueObject.CronSchedule
	if input["schedule"] != nil {
		schedule, err := valueObject.NewCronSchedule(input["schedule"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		schedulePtr = &schedule
	}

	var commandPtr *valueObject.UnixCommand
	if input["command"] != nil {
		command, err := valueObject.NewUnixCommand(input["command"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		commandPtr = &command
	}

	var commentPtr *valueObject.CronComment
	if input["comment"] != nil {
		comment, err := valueObject.NewCronComment(input["comment"])
		if err != nil {
			return NewServiceOutput(UserError, err.Error())
		}
		commentPtr = &comment
	}

	dto := dto.NewUpdateCron(id, schedulePtr, commandPtr, commentPtr)

	cmdRepo, err := cronInfra.NewCronCmdRepo()
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	err = useCase.UpdateCron(service.queryRepo, cmdRepo, dto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "CronUpdated")
}

func (service *CronService) Delete(input map[string]interface{}) ServiceOutput {
	id, err := valueObject.NewCronId(input["id"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	cmdRepo, err := cronInfra.NewCronCmdRepo()
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	dto := dto.NewDeleteCron(&id, nil)

	err = useCase.DeleteCron(service.queryRepo, cmdRepo, dto)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, "CronDeleted")
}

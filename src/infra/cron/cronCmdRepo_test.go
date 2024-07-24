package cronInfra

import (
	"testing"

	testHelpers "github.com/speedianet/os/src/devUtils"
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
)

func TestCronCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()
	cronCmdRepo, err := NewCronCmdRepo()
	if err != nil {
		t.Errorf("UnexpectedError: %v", err)
	}

	t.Run("CreateCron", func(t *testing.T) {
		schedule, err := valueObject.NewCronSchedule("* * * * *")
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}

		command, _ := valueObject.NewUnixCommand("echo \"cronTest\" >> crontab_log.txt")
		comment, _ := valueObject.NewCronComment("Test cron job")
		createCron := dto.NewCreateCron(schedule, command, &comment)

		err = cronCmdRepo.Create(createCron)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("UpdateCron", func(t *testing.T) {
		id, _ := valueObject.NewCronId(1)
		schedule, _ := valueObject.NewCronSchedule("* * * * 0")
		command, _ := valueObject.NewUnixCommand("echo \"cronUpdateTest\" >> crontab_logs.txt")
		comment, _ := valueObject.NewCronComment("update test")
		updateCron := dto.NewUpdateCron(id, &schedule, &command, &comment)

		err = cronCmdRepo.Update(updateCron)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("DeleteCron", func(t *testing.T) {
		id, _ := valueObject.NewCronId(1)
		err = cronCmdRepo.Delete(id)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})
}

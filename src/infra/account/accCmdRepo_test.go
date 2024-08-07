package accountInfra

import (
	"os"
	"testing"

	testHelpers "github.com/speedianet/os/src/devUtils"
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
)

func createDummyUser() error {
	username := valueObject.NewUsernamePanic(os.Getenv("DUMMY_USER_NAME"))
	password := valueObject.NewPasswordPanic(os.Getenv("DUMMY_USER_PASS"))

	createUser := dto.CreateAccount{
		Username: username,
		Password: password,
	}

	accCmdRepo := AccCmdRepo{}
	err := accCmdRepo.Create(createUser)
	if err != nil {
		return err
	}

	return nil
}

func deleteDummyUser() error {
	accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

	accCmdRepo := AccCmdRepo{}
	err := accCmdRepo.Delete(accountId)
	if err != nil {
		return err
	}

	return nil
}

func resetDummyUser() {
	_ = createDummyUser()
	_ = deleteDummyUser()
	_ = createDummyUser()
}

func TestAccCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()

	t.Run("CreateValidAccount", func(t *testing.T) {
		err := createDummyUser()
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}
	})

	t.Run("CreateInvalidAccount", func(t *testing.T) {
		username := valueObject.NewUsernamePanic("root")
		password := valueObject.NewPasswordPanic("invalid")

		createUser := dto.CreateAccount{
			Username: username,
			Password: password,
		}

		accCmdRepo := AccCmdRepo{}
		err := accCmdRepo.Create(createUser)
		if err == nil {
			t.Error("ExpectingError")
		}
	})

	t.Run("DeleteValidAccount", func(t *testing.T) {
		_ = createDummyUser()

		err := deleteDummyUser()
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}

		_ = createDummyUser()
	})

	t.Run("UpdatePasswordValidAccount", func(t *testing.T) {
		resetDummyUser()

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))
		newPassword := valueObject.NewPasswordPanic("newPassword")

		accCmdRepo := AccCmdRepo{}
		err := accCmdRepo.UpdatePassword(accountId, newPassword)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}

		resetDummyUser()
	})

	t.Run("UpdateApiKeyValidAccount", func(t *testing.T) {
		resetDummyUser()

		accountId := valueObject.NewAccountIdPanic(os.Getenv("DUMMY_USER_ID"))

		accCmdRepo := AccCmdRepo{}
		_, err := accCmdRepo.UpdateApiKey(accountId)
		if err != nil {
			t.Errorf("UnexpectedError: %v", err)
		}

		resetDummyUser()
	})
}

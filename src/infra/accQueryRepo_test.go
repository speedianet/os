package infra

import (
	"os"
	"testing"

	testHelpers "github.com/speedianet/sam/src/devUtils"
	"github.com/speedianet/sam/src/domain/valueObject"
)

func TestAccQueryRepo(t *testing.T) {
	testHelpers.LoadEnvVars()

	t.Run("GetValidAccountByUsername", func(t *testing.T) {
		username := valueObject.NewUsernamePanic(os.Getenv("DUMMY_USER_NAME"))

		authQueryRepo := AccQueryRepo{}
		_, err := authQueryRepo.GetByUsername(username)
		if err != nil {
			t.Error("UnexpectedError")
		}
	})

	t.Run("GetValidAccountById", func(t *testing.T) {
		userId := valueObject.NewUserIdFromStringPanic(os.Getenv("DUMMY_USER_ID"))

		authQueryRepo := AccQueryRepo{}
		_, err := authQueryRepo.GetById(userId)
		if err != nil {
			t.Error("UnexpectedError")
		}
	})

	t.Run("GetInvalidAccount", func(t *testing.T) {
		username := valueObject.NewUsernamePanic("invalid")

		authQueryRepo := AccQueryRepo{}

		_, err := authQueryRepo.GetByUsername(username)
		if err == nil {
			t.Error("ExpectingError")
		}
	})
}

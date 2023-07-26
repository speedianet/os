package infra

import (
	"os"
	"testing"

	testHelpers "github.com/speedianet/sam/src/devUtils"
	"github.com/speedianet/sam/src/domain/entity"
	"github.com/speedianet/sam/src/domain/valueObject"
	servicesInfra "github.com/speedianet/sam/src/infra/services"
)

func TestRuntimeCmdRepo(t *testing.T) {
	t.Skip("SkipRuntimeCmdRepoTest")
	testHelpers.LoadEnvVars()

	servicesInfra.Install(
		valueObject.NewServiceNamePanic("openlitespeed"),
		nil,
	)

	t.Run("UpdatePhpVersion", func(t *testing.T) {
		err := RuntimeCmdRepo{}.UpdatePhpVersion(
			valueObject.NewFqdnPanic(os.Getenv("VIRTUAL_HOST")),
			valueObject.NewPhpVersionPanic("8.1"),
		)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("UpdatePhpSettings", func(t *testing.T) {
		err := RuntimeCmdRepo{}.UpdatePhpSettings(
			valueObject.NewFqdnPanic(os.Getenv("VIRTUAL_HOST")),
			[]entity.PhpSetting{
				entity.NewPhpSetting(
					valueObject.NewPhpSettingNamePanic("display_errors"),
					valueObject.NewPhpSettingValuePanic("Off"),
					nil,
				),
			},
		)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("UpdatePhpModules", func(t *testing.T) {
		err := RuntimeCmdRepo{}.UpdatePhpModules(
			valueObject.NewFqdnPanic(os.Getenv("VIRTUAL_HOST")),
			[]entity.PhpModule{
				entity.NewPhpModule(
					valueObject.NewPhpModuleNamePanic("ioncube"),
					true,
				),
			},
		)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		err = RuntimeCmdRepo{}.UpdatePhpModules(
			valueObject.NewFqdnPanic(os.Getenv("VIRTUAL_HOST")),
			[]entity.PhpModule{
				entity.NewPhpModule(
					valueObject.NewPhpModuleNamePanic("ioncube"),
					false,
				),
			},
		)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}

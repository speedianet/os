package shared

import (
	"errors"

	"github.com/speedianet/sam/src/domain/valueObject"
	"github.com/speedianet/sam/src/infra"
)

func CheckServices(serviceNameStr string) error {
	servicesQueryRepo := infra.ServicesQueryRepo{}

	serviceName, err := valueObject.NewServiceName(serviceNameStr)

	currentSvcStatus, err := servicesQueryRepo.GetByName(serviceName)
	if err != nil {
		return err
	}

	var serviceErrorMessage string

	isStopped := currentSvcStatus.Status.String() == "stopped"
	if isStopped {
		serviceErrorMessage = "ServiceStopped"
	}
	isUninstalled := currentSvcStatus.Status.String() == "uninstalled"
	if isUninstalled {
		serviceErrorMessage = "ServiceNotInstalled"
	}
	shouldInstall := isStopped || isUninstalled
	if shouldInstall {
		return errors.New(serviceErrorMessage)
	}

	return nil
}

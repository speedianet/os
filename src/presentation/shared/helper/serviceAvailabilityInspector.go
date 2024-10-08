package sharedHelper

import (
	"github.com/goinfinite/os/src/domain/valueObject"
	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	servicesInfra "github.com/goinfinite/os/src/infra/services"
)

const ServiceUnavailableError = "ServiceUnavailable"

type ServiceAvailabilityInspector struct {
	servicesQueryRepo *servicesInfra.ServicesQueryRepo
}

func NewServiceAvailabilityInspector(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *ServiceAvailabilityInspector {
	return &ServiceAvailabilityInspector{
		servicesQueryRepo: servicesInfra.NewServicesQueryRepo(persistentDbSvc),
	}
}

func (inspector *ServiceAvailabilityInspector) IsAvailable(
	serviceName valueObject.ServiceName,
) bool {
	availableSvc, err := inspector.servicesQueryRepo.ReadByName(serviceName)
	if err != nil {
		return false
	}

	return availableSvc.Status.String() == "running"
}

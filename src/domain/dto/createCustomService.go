package dto

import "github.com/speedianet/os/src/domain/valueObject"

type CreateCustomService struct {
	Name              valueObject.ServiceName     `json:"name"`
	Type              valueObject.ServiceType     `json:"type"`
	Command           valueObject.UnixCommand     `json:"command"`
	Version           *valueObject.ServiceVersion `json:"version"`
	PortBindings      []valueObject.PortBinding   `json:"portBindings"`
	AutoCreateMapping bool                        `json:"autoCreateMapping"`
}

func NewCreateCustomService(
	name valueObject.ServiceName,
	serviceType valueObject.ServiceType,
	command valueObject.UnixCommand,
	version *valueObject.ServiceVersion,
	portBindings []valueObject.PortBinding,
	autoCreateMapping bool,
) CreateCustomService {
	return CreateCustomService{
		Name:              name,
		Type:              serviceType,
		Command:           command,
		Version:           version,
		PortBindings:      portBindings,
		AutoCreateMapping: autoCreateMapping,
	}
}

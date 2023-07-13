package infra

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/speedianet/sam/src/domain/entity"
	"github.com/speedianet/sam/src/domain/valueObject"
	"golang.org/x/exp/slices"

	"github.com/shirou/gopsutil/process"
)

type ServicesQueryRepo struct {
}

var KnownServices = []string{
	"nginx",
	"openlitespeed",
	"node",
	"python",
	"ruby",
	"java",
	"mysqld",
	"redis",
	"elasticsearch",
}

func (repo ServicesQueryRepo) runningServiceFactory() ([]entity.Service, error) {
	pids, err := process.Pids()
	if err != nil {
		return []entity.Service{}, errors.New("PidsUnavailable")
	}

	runningStatus, _ := valueObject.NewServiceStatus("running")

	var services []entity.Service
	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			continue
		}
		pidUint := uint32(pid)

		fullName, err := p.Name()
		if err != nil {
			continue
		}
		nameSlice := strings.Fields(fullName)
		if len(nameSlice) == 0 {
			continue
		}
		svcName, err := valueObject.NewServiceName(nameSlice[0])
		if err != nil {
			continue
		}
		if !slices.Contains(KnownServices, svcName.String()) {
			continue
		}

		uptime, err := p.CreateTime()
		if err != nil {
			continue
		}

		uptimeSeconds := time.Since(time.Unix(uptime/1000, 0)).Seconds()

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			continue
		}

		memPercent, err := p.MemoryPercent()
		if err != nil {
			continue
		}

		services = append(
			services,
			entity.NewService(
				svcName,
				runningStatus,
				&pidUint,
				&uptimeSeconds,
				&cpuPercent,
				&memPercent,
			),
		)
	}

	return services, nil
}

func (repo ServicesQueryRepo) Get() ([]entity.Service, error) {
	runningServices, err := repo.runningServiceFactory()
	if err != nil {
		return []entity.Service{}, err
	}

	var runningServicesNames []string
	for _, svc := range runningServices {
		runningServicesNames = append(runningServicesNames, svc.Name.String())
	}

	var notRunningServicesNames []string
	for _, svc := range KnownServices {
		if !slices.Contains(runningServicesNames, svc) {
			notRunningServicesNames = append(notRunningServicesNames, svc)
		}
	}

	var remainingServices []entity.Service
	confFilePath := "/speedia/supervisord.conf"
	for _, svc := range notRunningServicesNames {
		cmd := exec.Command(
			"awk",
			fmt.Sprintf("/%s/{found=1} END{if(!found) exit 1}", svc),
			confFilePath,
		)
		err := cmd.Run()

		svcName, _ := valueObject.NewServiceName(svc)
		svcStatus, _ := valueObject.NewServiceStatus("stopped")
		if err != nil {
			svcStatus, _ = valueObject.NewServiceStatus("uninstalled")
		}

		remainingServices = append(
			remainingServices,
			entity.NewService(
				svcName,
				svcStatus,
				nil,
				nil,
				nil,
				nil,
			),
		)
	}

	return append(runningServices, remainingServices...), nil
}

func (repo ServicesQueryRepo) GetByName(
	name valueObject.ServiceName,
) (entity.Service, error) {
	services, err := repo.Get()
	if err != nil {
		return entity.Service{}, err
	}

	for _, svc := range services {
		if svc.Name.String() == name.String() {
			return svc, nil
		}
	}

	return entity.Service{}, errors.New("ServiceNotFound")
}

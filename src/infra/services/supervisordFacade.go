package servicesInfra

import (
	"errors"
	"log"
	"os"

	"github.com/speedianet/sam/src/domain/valueObject"
	infraHelper "github.com/speedianet/sam/src/infra/helper"
)

type SupervisordFacade struct {
}

const supervisordCmd string = "/usr/bin/supervisord"

func (facade SupervisordFacade) Start(name valueObject.ServiceName) error {
	_, err := infraHelper.RunCmd(
		supervisordCmd,
		"ctl",
		"start",
		name.String(),
	)
	if err != nil {
		log.Printf("StartServiceError: %s", err)
		return errors.New("StartServiceError")
	}

	return nil
}

func (facade SupervisordFacade) Stop(name valueObject.ServiceName) error {
	_, err := infraHelper.RunCmd(
		supervisordCmd,
		"ctl",
		"stop",
		name.String(),
	)
	if err != nil {
		log.Printf("StopServiceError: %s", err)
		return errors.New("StopServiceError")
	}

	return nil
}

func (facade SupervisordFacade) Reload() error {
	_, err := infraHelper.RunCmd(
		supervisordCmd,
		"ctl",
		"reload",
	)
	if err != nil {
		log.Printf("ReloadSupervisorError: %s", err)
		return errors.New("ReloadSupervisorError")
	}

	return nil
}

func (facade SupervisordFacade) AddConf(svcName string, svcCmd string) error {
	supervisorConf := `
[program:` + svcName + `]
command=` + svcCmd + `
user=root
directory=/speedia
autostart=true
autorestart=true
startretries=3
startsecs=5
stdout_logfile=/dev/stdout
stdout_logfile_maxbytes=0
stderr_logfile=/dev/stderr
stderr_logfile_maxbytes=0
`

	f, err := os.OpenFile("/speedia/supervisord.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("OpenSupervisorConfError: %s", err)
		return errors.New("OpenSupervisorConfError")
	}
	defer f.Close()

	if _, err := f.WriteString(supervisorConf); err != nil {
		log.Printf("WriteSupervisorConfError: %s", err)
		return errors.New("WriteSupervisorConfError")
	}

	return nil
}

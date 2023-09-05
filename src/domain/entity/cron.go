package entity

import "github.com/speedianet/sam/src/domain/valueObject"

type Cron struct {
	Id       valueObject.CronId       `json:"id"`
	Schedule valueObject.CronSchedule `json:"schedule"`
	Command  valueObject.UnixCommand  `json:"command"`
	Comment  *valueObject.CronComment `json:"comment"`
}

func NewCron(
	id valueObject.CronId,
	schedule valueObject.CronSchedule,
	command valueObject.UnixCommand,
	comment *valueObject.CronComment,
) Cron {
	return Cron{
		Id:       id,
		Schedule: schedule,
		Command:  command,
		Comment:  comment,
	}
}

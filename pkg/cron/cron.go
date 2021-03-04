package cron

import (
	"github.com/robfig/cron"
)

type Schedule struct {
	ScheduleCron *cron.Cron
}

func NewCron() *cron.Cron  {
	return cron.New()

}


func (s *Schedule)AddJob(options ...func())  {

	for _, option := range options {
		s.ScheduleCron.AddFunc("",option)
	}
}

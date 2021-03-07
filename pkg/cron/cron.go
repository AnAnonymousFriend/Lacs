package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type Schedule struct {
	ScheduleCron *cron.Cron
}

type ScheduleError error

// 声明一个新的调度器
func NewSchedule() *Schedule  {
	return  &Schedule{
		ScheduleCron:cron.New(),
	}
}

// 添加任务
func (s *Schedule) AddScheduleJob(options ...func()) *Schedule  {
	for _, option := range options {
		ScheduleError := s.ScheduleCron.AddFunc("* * * * * *", option)
		if ScheduleError != nil {
			fmt.Println(ScheduleError)
		}
	}
	return s
}

// 启动调度器
func (s *Schedule) ScheduleRun(minutes int32) {
	s.ScheduleCron.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			//println("执行任务")
			t1.Reset(time.Second * 10)
		}
	}
	//t1 := time.NewTimer(time.Duration(minutes) * time.Minute)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Duration(minutes) * time.Minute)
	//	}
	//}
}

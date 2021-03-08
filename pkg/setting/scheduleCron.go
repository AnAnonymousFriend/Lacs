package setting

import (
	c "Lacs/pkg/cron"
	s "Lacs/server"
)

func ScheduleSetup()  {
	job := c.NewSchedule()
	job.AddScheduleJob(s.PingDevices)
	go job.ScheduleRun(1)
}


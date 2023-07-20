package src

import (
	"github.com/robfig/cron/v3"

	"github.com/maulanar/kms/app"
)

func Scheduler() *schedulerUtil {
	if scheduler == nil {
		scheduler = &schedulerUtil{}
		scheduler.Configure()
		scheduler.isConfigured = true
	}
	return scheduler
}

var scheduler *schedulerUtil

type schedulerUtil struct {
	isConfigured bool
}

func (s *schedulerUtil) Configure() {
	c := cron.New()

	// add scheduler func here, for example :
	c.AddFunc("CRON_TZ=Asia/Jakarta * * * * *", app.RemoveExpiredToken)

	c.Start()
}

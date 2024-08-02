package schedule

import (
	"NaimBiswas/email-reminder/schedules/daily"
	"NaimBiswas/email-reminder/schedules/hourly"
	"NaimBiswas/email-reminder/schedules/minute"
	"NaimBiswas/email-reminder/schedules/monthly"

	"github.com/robfig/cron/v3"
)

func Minute(c *cron.Cron) {
	minute.SetUpMinuteCron(c)
}

func Hourly(c *cron.Cron) {
	hourly.SetUpHourlyCron(c)
}

func Daily(c *cron.Cron) {
	daily.SetUpDailyCron(c)
}

func Monthly(c *cron.Cron) {
	monthly.SetUpMonthlyJob(c)
}

package schedule

import (
	"NaimBiswas/email-reminder/schedules/daily"
	"NaimBiswas/email-reminder/schedules/hourly"
	"NaimBiswas/email-reminder/schedules/minute"
	"NaimBiswas/email-reminder/schedules/monthly"

	"github.com/robfig/cron/v3"
)

func Minute(c *cron.Cron, expression string) {
	minute.SetUpMinuteCron(c, expression)
}

func Hourly(c *cron.Cron, expression string) {
	hourly.SetUpHourlyCron(c, expression)
}

func Daily(c *cron.Cron, expression string) {
	daily.SetUpDailyCron(c, expression)
}

func Monthly(c *cron.Cron,  expression string) {
	monthly.SetUpMonthlyJob(c, expression)
}

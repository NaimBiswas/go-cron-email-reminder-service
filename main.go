package main

import (
	schedule "NaimBiswas/email-reminder/schedules"
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Sprintf("%s", "❤️‍🔥❤️‍🔥 Welcome for Schedule Cron Services❤️‍🔥❤️‍🔥")

	cron := cron.New()

	schedule.Minute(cron)
	fmt.Println("Minute Job has been started. Job Will Execute at every minute")

	schedule.Hourly(cron)
	fmt.Println("Hourly Job has been started. Job Will Execute at every hour")

	schedule.Daily(cron)
	fmt.Println("Daily Job has been started. Job Will Execute at every day 12.00 AM")

	schedule.Monthly(cron)
	fmt.Println("Monthly Job has been started. Job Will Execute at every month 5 11.30 AM")

	cron.Start()

	// Keep the program running
	select {}
}

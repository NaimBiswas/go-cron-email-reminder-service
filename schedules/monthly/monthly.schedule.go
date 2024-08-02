package monthly

import (
	"NaimBiswas/email-reminder/services"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func SetUpMonthlyJob(c *cron.Cron) {
	_, err := c.AddFunc("30 11 5 * *", func() {
		fmt.Println("Minute Job start executing at ==>", time.Now().Format("02/01/2006 03:04:05 PM MST"))
		services.SendEmailService([]string{}, "Niraj Shah, Rana Mazen")
	})
	if err != nil {
		log.Fatal(err)
	}
}

package minute

import (
	"NaimBiswas/email-reminder/services"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func SetUpMinuteCron(c *cron.Cron) {
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("Minute Job start executing at ==>", time.Now().Format("02/01/2006 03:04:05 PM MST"))
		services.SendEmailService([]string{"naimbiswas221@gmail.com", "nayeembiswas2@gmail.com"}, "Niraj Shah, Rana Mazen")
	})
	if err != nil {
		log.Fatal(err)
	}
}

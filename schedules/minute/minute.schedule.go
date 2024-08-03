package minute

import (
	"NaimBiswas/email-reminder/schedules/common"
	"NaimBiswas/email-reminder/services"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func SetUpMinuteCron(c *cron.Cron, expression string) {
	_, err := c.AddFunc(expression, func() {
		fmt.Println("Minute Job start executing at ==>", time.Now().Format("02/01/2006 03:04:05 PM MST"))
		services.SendEmailService(common.GetHREmails(os.Getenv("HR_EMAILS")), os.Getenv("HR_NAMES"))
	})
	if err != nil {
		log.Fatal(err)
	}
}

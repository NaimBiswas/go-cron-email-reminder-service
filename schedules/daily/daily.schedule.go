package daily

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func SetUpDailyCron(c *cron.Cron) {
	_, err := c.AddFunc("0 0 * * *", func() {
		fmt.Println("Daily Job is running")
	})
	if err != nil {
		log.Fatal(err)
	}
}

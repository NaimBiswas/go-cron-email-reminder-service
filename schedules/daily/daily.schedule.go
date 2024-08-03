package daily

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func SetUpDailyCron(c *cron.Cron, expression string) {
	_, err := c.AddFunc(expression, func() {
		fmt.Println("Daily Job is running")
	})
	if err != nil {
		log.Fatal(err)
	}
}

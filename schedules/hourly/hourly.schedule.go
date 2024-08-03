package hourly

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func SetUpHourlyCron(c *cron.Cron, expression string) {
	_, err := c.AddFunc(expression, func() {
		fmt.Println("Hourly Job is running")
	})
	if err != nil {
		log.Fatal(err)
	}
}

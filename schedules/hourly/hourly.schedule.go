package hourly

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func SetUpHourlyCron(c *cron.Cron) {
	_, err := c.AddFunc("0 * * * *", func() {
		fmt.Println("Hourly Job is running")
	})
	if err != nil {
		log.Fatal(err)
	}
}

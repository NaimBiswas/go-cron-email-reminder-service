package common

import (
	"fmt"
	"strings"
)

func GetHREmails(emails string) []string {
	emailIds := strings.Split(emails, ",")
	fmt.Println(emailIds)
	return emailIds
}
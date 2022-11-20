package internal

import (
	"log"
	"time"
)

// personioDateStr2Date converts string date to time.Time
func DateStr2Date(dateString string) (time.Time, error) {
	layout := "2006-01-02T15:04:05-0700"

	dateResult, err := time.Parse(layout, dateString)
	if err != nil {
		log.Print("Invalid date string!", err)
		return time.Time{}, err
	}

	return dateResult, nil
}


func InitName2MailMap(data []map[string]string) map[string]string {
	name2email := make(map[string]string, len(data))

	for _, elem := range data {
		name2email[elem["name"]] = elem["email"]
	}
	return name2email
}
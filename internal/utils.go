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

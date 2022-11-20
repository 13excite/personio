package internal

import (
	//"fmt"

	"testing"
	"time"
)

func TestDateStr2Date(t *testing.T) {

	testStringDate := "2022-11-11T10:11:11-0000"
	expectedDate := time.Date(2022, 11, 11, 10, 11, 11, 0, time.UTC)

	resultDate, err := DateStr2Date(testStringDate)

	if err != nil {
		t.Errorf("Got error from DateStr2Date")
		return
	}

	if !expectedDate.Equal(resultDate.UTC()) {
		t.Errorf("DateStr2Date returns incorrect date. Expected: %s. Got: %s", expectedDate, resultDate.UTC())

	}
}

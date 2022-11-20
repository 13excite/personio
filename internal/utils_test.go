package internal

import (
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

func TestInitName2MailMap(t *testing.T) {
	testData := []struct {
		name  string
		email string
	}{
		{"Jim Jim", "jim@example.com"},
		{"Jack Jack", "jack@example.com"},
	}
	var testSlice []map[string]string

	for _, testD := range testData {
		testSlice = append(testSlice, map[string]string{
			"name":  testD.name,
			"email": testD.email,
		})
	}

	resultMap := InitName2MailMap(testSlice)

	for _, testD := range testData {
		if val, ok := resultMap[testD.name]; ok {
			if val != testD.email {
				t.Errorf("Key contains unexpected values. Expected: %s, Got: %s", testD.email, val)
			}
		} else {
			t.Errorf("Expected key %s not found in map", testD.name)
		}
	}

}

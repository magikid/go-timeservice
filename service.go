package main

import (
	"fmt"
	"log"
	"time"
)

type TimeRequest struct {
	format   string
	timezone string
}

type TimeResponse struct {
	Time     string
	Timezone string
}

func timeResponse(t TimeRequest) TimeResponse {
	timeNow := time.Now()

	timeLocation, err := time.LoadLocation(t.timezone)
	if err != nil {
		log.Println(fmt.Sprintf("Invalid timezone: %s", t.timezone))
		timeLocation = time.FixedZone("UTC", 0)
	}

	timeWithLoc := timeNow.In(timeLocation)

	var timeString string
	if t.format == "timestamp" {
		timeString = fmt.Sprintf("%d", timeWithLoc.Unix())
	} else {
		timeString = timeWithLoc.Format(time.UnixDate)
	}

	response := TimeResponse{timeString, timeLocation.String()}
	return response
}

func (tr TimeResponse) String() string {
	return fmt.Sprintf("%s, %s", tr.Time, tr.Timezone)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const timeLayout = "15:04"

func main() {
	timezone := flag.String("tz", "America/New_York", "Timezone to convert to. List of timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
	inputTime := flag.String("time", "", "Time to convert (default \"Current time\"), format HH:mm")

	flag.Parse()

	location, err := time.LoadLocation(*timezone)
	if err != nil {
		log.Fatal("Could not load timezone: ", err)
	}

	var timeToConvert time.Time
	if *inputTime == "" {
		timeToConvert = time.Now()
	} else {
		parsedInputTime, err := time.Parse(timeLayout, *inputTime)
		if err != nil {
			log.Fatal("Could not parse time: ", err)
		}
		now := time.Now()
		timeToConvert = time.Date(now.Year(), now.Month(), now.Day(), parsedInputTime.Hour(), parsedInputTime.Minute(), 0, now.Nanosecond(), now.Location())
	}

	convertedTime := timeToConvert.In(location)
	fmt.Printf("%s is %s in %s", timeToConvert.Format(timeLayout), convertedTime.Format(timeLayout), *timezone)
}

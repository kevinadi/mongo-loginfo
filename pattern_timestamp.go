package main

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

type Res_LogTimes struct {
	log_start    time.Time
	log_end      time.Time
	log_duration time.Duration
	log_length   int
}

var res_logtimes = new(Res_LogTimes)

const Timestamp_pattern = "2006-01-02T15:04:05-0700"

var Timestamp_regex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}[-+]\d{4}$`)

func parse_timestamp(val string) time.Time {
	var t time.Time
	var err error
	if Timestamp_regex.FindString(val) != "" {
		t, err = time.Parse(Timestamp_pattern, val)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return t
}

func Matcher_timestamp(line <-chan string, result chan<- *Res_LogTimes, wg_main *sync.WaitGroup) {
	var cur string
	var linecount int
	var time_start, time_end time.Time

	for val := range line {
		linecount += 1
		if linecount == 1 {
			time_start = parse_timestamp(val)
		}
		if Timestamp_regex.FindString(val) != "" {
			cur = val
		}
	}

	res_logtimes.log_start = time_start
	res_logtimes.log_end = parse_timestamp(cur)
	res_logtimes.log_duration = time_end.Sub(time_start)
	res_logtimes.log_length = linecount

	result <- res_logtimes
	close(result)
	wg_main.Done()
}

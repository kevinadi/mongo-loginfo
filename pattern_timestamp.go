package main

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

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

func Matcher_timestamp(line chan string, ts *time.Time, wg_main *sync.WaitGroup) {
	var cur string
	for val := range line {
		if Timestamp_regex.FindString(val) != "" {
			cur = val
		}
	}
	t, err := time.Parse(Timestamp_pattern, cur)
	if err != nil {
		fmt.Println(err)
	}
	*ts = t
	wg_main.Done()
}

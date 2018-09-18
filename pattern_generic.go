package main

import (
	"fmt"
	"regexp"
	"time"
)

type DatePattern struct {
	title     string
	regex     *regexp.Regexp
	matchfunc func(string) time.Time
}

var pattern_date = DatePattern{
	"timestamp",
	regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}-\d{4}$`),
	func(match string) time.Time {
		pattern := "2006-01-02T15:04:05-0700"
		var t time.Time
		var err error
		if len(match) > 1 {
			t, err = time.Parse(pattern, match)
			if err != nil {
				fmt.Println(err)
			}
		}
		return t
	},
}

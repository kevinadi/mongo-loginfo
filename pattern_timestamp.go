package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Res_LogTimes struct {
	log_start    time.Time
	log_end      time.Time
	log_duration string
	log_length   int
}

var res_logtimes = new(Res_LogTimes)

const Timestamp_pattern = "2006-01-02T15:04:05-0700"
const Timestamp_pattern_Z = "2006-01-02T15:04:05Z"

var Timestamp_regex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}(Z|[-+]\d{4})$`)

func parse_timestamp(val string) time.Time {
	var t time.Time
	var err error
	pat := Timestamp_pattern
	if Timestamp_regex.FindString(val) != "" {
		if strings.HasSuffix(val, "Z") {
			pat = Timestamp_pattern_Z
		}
		t, err = time.Parse(pat, val)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return t
}

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65
func duration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}

func Matcher_timestamp(line <-chan string, result chan<- *Res_LogTimes, wg_main *sync.WaitGroup) {
	var cur string
	var linecount int
	var time_start time.Time

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
	res_logtimes.log_duration = duration(res_logtimes.log_end.Sub(time_start))
	res_logtimes.log_length = linecount

	result <- res_logtimes
	close(result)
	wg_main.Done()
}

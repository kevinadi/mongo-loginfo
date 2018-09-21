package main

import (
	"sync"
	"testing"
	"time"
)

func Test_parsetimestamp_1(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799+1100"
	parsed := parse_timestamp(timestr)

	tz := time.FixedZone("+11", +11*60*60)
	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, tz)
	if !expect.Equal(parsed) {
		t.Error("parsed time is", parsed, "expecting", expect)
	}
}

func Test_parsetimestamp_2(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799-1100"
	parsed := parse_timestamp(timestr)

	tz := time.FixedZone("-11", -11*60*60)
	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, tz)
	if !expect.Equal(parsed) {
		t.Error("parsed time is", parsed, "expecting", expect)
	}
}

func Test_parsetimestamp_3(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799Z"
	parsed := parse_timestamp(timestr)

	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, time.UTC)
	if !expect.Equal(parsed) {
		t.Error("parsed time is", parsed, "expecting", expect)
	}
}

func Test_timestamp_1(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan string)
	outch := make(chan *Res_LogTimes)
	go Matcher_timestamp(ch, outch, &wg)
	wg.Add(1)

	start_time := "2018-03-13T10:15:44.799+1100"
	end_time := "2018-03-14T10:15:44.799+1100"
	ch <- start_time
	ch <- end_time
	close(ch)

	res := <-outch
	wg.Wait()

	if !res.log_start.Equal(parse_timestamp(start_time)) {
		t.Error("time_start is", res.log_start, "expecting", start_time)
	}
	if !res.log_end.Equal(parse_timestamp(end_time)) {
		t.Error("time_end is", res.log_end, "expecting", end_time)
	}
	if res.log_duration != "1 day" {
		t.Error("duration is", res.log_duration, "expecting 1 day")
	}
	if res.log_length != 2 {
		t.Error("length is", res.log_length, "expecting 2")
	}
}

func Test_timestamp_2(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan string)
	outch := make(chan *Res_LogTimes)
	go Matcher_timestamp(ch, outch, &wg)
	wg.Add(1)

	start_time := "2018-03-13T10:15:44.799-1100"
	end_time := "2018-03-14T10:15:44.799-1100"
	ch <- start_time
	ch <- end_time
	close(ch)

	res := <-outch
	wg.Wait()

	if !res.log_start.Equal(parse_timestamp(start_time)) {
		t.Error("time_start is", res.log_start, "expecting", start_time)
	}
	if !res.log_end.Equal(parse_timestamp(end_time)) {
		t.Error("time_end is", res.log_end, "expecting", end_time)
	}
	if res.log_duration != "1 day" {
		t.Error("duration is", res.log_duration, "expecting 1 day")
	}
	if res.log_length != 2 {
		t.Error("length is", res.log_length, "expecting 2")
	}
}

func Test_timestamp_3(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan string)
	outch := make(chan *Res_LogTimes)
	go Matcher_timestamp(ch, outch, &wg)
	wg.Add(1)

	start_time := "2018-03-13T10:15:44.799Z"
	end_time := "2018-03-14T10:15:44.799Z"
	ch <- start_time
	ch <- end_time
	close(ch)

	res := <-outch
	wg.Wait()

	if !res.log_start.Equal(parse_timestamp(start_time)) {
		t.Error("time_start is", res.log_start, "expecting", start_time)
	}
	if !res.log_end.Equal(parse_timestamp(end_time)) {
		t.Error("time_end is", res.log_end, "expecting", end_time)
	}
	if res.log_duration != "1 day" {
		t.Error("duration is", res.log_duration, "expecting 1 day")
	}
	if res.log_length != 2 {
		t.Error("length is", res.log_length, "expecting 2")
	}
}

func Test_timestamp_4(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan string)
	outch := make(chan *Res_LogTimes)
	go Matcher_timestamp(ch, outch, &wg)
	wg.Add(1)

	start_time := "2018-03-13T10:15:44.799Z"
	end_time := "2018-03-14T10:15:44.799Z"
	ch <- start_time
	ch <- end_time
	ch <- "***"
	ch <- ""
	close(ch)

	res := <-outch
	wg.Wait()

	if !res.log_start.Equal(parse_timestamp(start_time)) {
		t.Error("time_start is", res.log_start, "expecting", start_time)
	}
	if !res.log_end.Equal(parse_timestamp(end_time)) {
		t.Error("time_end is", res.log_end, "expecting", end_time)
	}
	if res.log_duration != "1 day" {
		t.Error("duration is", res.log_duration, "expecting 1 day")
	}
	if res.log_length != 4 {
		t.Error("length is", res.log_length, "expecting 2")
	}
}

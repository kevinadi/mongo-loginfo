package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_parsetimestamp_1(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799+1100"
	parsed := parse_timestamp(timestr)

	tz := time.FixedZone("+11", +11*60*60)
	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, tz)
	assert.True(t, expect.Equal(parsed))
}

func Test_parsetimestamp_2(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799-1100"
	parsed := parse_timestamp(timestr)

	tz := time.FixedZone("-11", -11*60*60)
	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, tz)
	assert.True(t, expect.Equal(parsed))
}

func Test_parsetimestamp_3(t *testing.T) {
	timestr := "2018-03-13T10:15:44.799Z"
	parsed := parse_timestamp(timestr)

	expect := time.Date(2018, 3, 13, 10, 15, 44, 799*1e6, time.UTC)
	assert.True(t, expect.Equal(parsed))
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

	assert.True(t, res.log_start.Equal(parse_timestamp(start_time)))
	assert.True(t, res.log_end.Equal(parse_timestamp(end_time)))
	assert.Equal(t, "1 day", res.log_duration)
	assert.Equal(t, 2, res.log_length)
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

	assert.True(t, res.log_start.Equal(parse_timestamp(start_time)))
	assert.True(t, res.log_end.Equal(parse_timestamp(end_time)))
	assert.Equal(t, "1 day", res.log_duration)
	assert.Equal(t, 2, res.log_length)
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

	assert.True(t, res.log_start.Equal(parse_timestamp(start_time)))
	assert.True(t, res.log_end.Equal(parse_timestamp(end_time)))
	assert.Equal(t, "1 day", res.log_duration)
	assert.Equal(t, 2, res.log_length)
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

	assert.True(t, res.log_start.Equal(parse_timestamp(start_time)))
	assert.True(t, res.log_end.Equal(parse_timestamp(end_time)))
	assert.Equal(t, "1 day", res.log_duration)
	assert.Equal(t, 4, res.log_length)
}

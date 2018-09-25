package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_conn_automation(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_automation(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-04-06T15:44:27.119-0500 I ACCESS   [conn2598] Successfully authenticated as principal mms-automation on admin"
	close(ch)
	wg.Wait()

	assert.Equal(t, res, "mms-automation")
}

func Test_conn_monitoring(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_monitoring(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-09-24T01:20:01.896+0000 I ACCESS   [conn90245] Successfully authenticated as principal mms-monitoring-agent on admin"
	close(ch)
	wg.Wait()

	assert.Equal(t, "mms-monitoring", res)
}

func Test_conn_db_version(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_db_version(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-09-24T01:19:50.845+0000 I CONTROL  [conn1597934] db version v3.6.4"
	close(ch)
	wg.Wait()

	assert.Equal(t, "3.6.4", res)
}

func Test_conn_host(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_host(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-09-24T01:19:50.844+0000 I CONTROL  [conn1597934] pid=30930 port=27017 64-bit host=myhost.local"
	close(ch)
	wg.Wait()

	assert.Equal(t, "myhost.local", res)
}

func Test_conn_port(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_port(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-09-24T01:19:50.844+0000 I CONTROL  [conn1597934] pid=30930 port=27017 64-bit host=myhost.local"
	close(ch)
	wg.Wait()

	assert.Equal(t, "27017", res)
}


package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main_restarts(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_restarts(&res, ch, &wg)
	wg.Add(1)

	ch <- "2018-03-13T10:14:44.799+1100 I CONTROL  [main] ***** SERVER RESTARTED *****"
	ch <- "2018-03-13T10:15:44.799+1100 I CONTROL  [main] ***** SERVER RESTARTED *****"
	close(ch)
	wg.Wait()

	assert.Equal(t, "2", res)
}

func Test_main_fasserts(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_fasserts(&res, ch, &wg)
	wg.Add(1)

	ch <- "2019-06-15T04:36:45.898+0000 F -        [conn53725] Fatal Assertion 28559 at src/mongo/db/storage/wiredtiger/wiredtiger_util.cpp 64"
	ch <- "2019-06-15T04:36:45.898+0000 F -        [thread53716] Fatal Assertion 28558 at src/mongo/db/storage/wiredtiger/wiredtiger_util.cpp 366"
	ch <- "2019-06-15T04:36:45.898+0000 F -        [conn53721] Fatal Assertion 28559 at src/mongo/db/storage/wiredtiger/wiredtiger_util.cpp 64"
	ch <- "2019-06-15T04:36:45.902+0000 F -        [conn53725]"
	ch <- ""
	ch <- "***aborting after fassert() failure"

	close(ch)
	wg.Wait()

	assert.Equal(t, "1", res)
}

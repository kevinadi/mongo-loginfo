package main

import (
	"sync"
	"testing"
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

	if res != "mms-automation" {
		t.Error("Automation is", res, "expecting mms-automation")
	}
}

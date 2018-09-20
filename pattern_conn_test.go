package main

import (
	"sync"
	"testing"
)

func Setup_conn_test(wg *sync.WaitGroup) {
	ch := make(chan string)

	go func() {
		ch <- "2018-04-06T15:44:27.119-0500 I ACCESS   [conn2598] Successfully authenticated as principal mms-automation on admin"
		close(ch)
	}()

	go Matcher(func_array_conn, ch, output, wg)
	wg.Add(1)
}

func Test_conn_automation(t *testing.T) {
	var wg sync.WaitGroup
	Setup_conn_test(&wg)

	wg.Wait()

	if res_conn.automation != "mms-automation" {
		t.Error("Automation is", res_conn.automation, "expecting mms-automation")
	}
}

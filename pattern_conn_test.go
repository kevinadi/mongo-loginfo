package main

import (
	"sync"
	"testing"
)

func TestPatternConn(t *testing.T) {
	var wg_main sync.WaitGroup
	ch := make(chan string)

	go Matcher(func_array_conn, ch, output, &wg_main)
	wg_main.Add(1)

	ch <- "2018-04-06T15:44:27.119-0500 I ACCESS   [conn2598] Successfully authenticated as principal mms-automation on admin"
	close(ch)
	wg_main.Wait()

	if res_conn.automation != "mms-automation" {
		t.Error("Automation is", res_conn.automation, "expecting mms-automation")
	}
}

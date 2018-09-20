package main

import (
	"sync"
	"testing"
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

	if res != "2" {
		t.Error("restart is", res, "expecting 2")
	}
}
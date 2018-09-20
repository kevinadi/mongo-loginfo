package main

import (
	"sync"
	"testing"
)

func Setup_main_test(wg *sync.WaitGroup) {
	ch := make(chan string)

	go func() {
		ch <- "2018-03-13T10:14:44.799+1100 I CONTROL  [main] ***** SERVER RESTARTED *****"
		ch <- "2018-03-13T10:15:44.799+1100 I CONTROL  [main] ***** SERVER RESTARTED *****"
		close(ch)
	}()

	go Matcher(func_array_main, ch, output, wg)
	wg.Add(1)
}

func Test_main_restart(t *testing.T) {
	var wg sync.WaitGroup
	r := res_main
	Setup_main_test(&wg)

	wg.Wait()

	if r.restarts != 2 {
		t.Error("Restart is", r.restarts, "expecting 2")
	}
}

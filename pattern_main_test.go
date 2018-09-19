package main

import (
	"sync"
	"testing"
)

func TestPatternMain(t *testing.T) {
	var wg_main sync.WaitGroup
	ch := make(chan string)

	go Matcher(func_array_main, ch, output, &wg_main)
	wg_main.Add(1)

	ch <- "xxx xxx xxx [main] *** SERVER RESTARTED ***"
	ch <- "*** SERVER RESTARTED ***"
	close(ch)
	wg_main.Wait()

	if res_main.restarts != 2 {
		t.Error("Restart is", res_main.restarts, "expecting 2")
	}
}

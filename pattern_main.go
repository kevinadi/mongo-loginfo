package main

import (
	"sync"
)

type Res_Main struct {
	restarts string
	fasserts string
}

var Matcher_restarts = RegexMatcher_count(`SERVER RESTARTED`)
var Matcher_fasserts = RegexMatcher_count(`\*\*\*aborting after fassert\(\)\ failure`)

func MatcherGroup_main(line <-chan string, result chan<- *Res_Main, wg_main *sync.WaitGroup) {
	var res_main = new(Res_Main)
	var Matchers_main = []MatcherType{
		MatcherType{Matcher_restarts, &res_main.restarts},
		MatcherType{Matcher_fasserts, &res_main.fasserts},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go RegexMatchers(Matchers_main, line, &wg)
	wg.Wait()
	result <- res_main
	close(result)
	wg_main.Done()
}

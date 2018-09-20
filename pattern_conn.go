package main

import "sync"

type Res_Conn struct {
	automation string
}

var Matcher_automation = RegexMatcher_string(`Successfully authenticated as principal (mms-automation)`)

func MatcherGroup_conn(line <-chan string, result chan<- *Res_Conn, wg_main *sync.WaitGroup) {
	var res_conn = new(Res_Conn)
	var Matchers_conn = []MatcherType{
		MatcherType{Matcher_automation, &res_conn.automation},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go RegexMatchers(Matchers_conn, line, &wg)
	wg.Wait()
	result <- res_conn
	close(result)
	wg_main.Done()
}

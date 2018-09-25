package main

import "sync"

type Res_Conn struct {
	automation string
	monitoring string
}

func (r *Res_Conn) Init() {
	r.automation = "-"
	r.monitoring = "-"
}

var Matcher_automation = RegexMatcher_string(`Successfully authenticated as principal (mms-automation)`)
var Matcher_monitoring = RegexMatcher_string(`Successfully authenticated as principal (mms-monitoring)-agent`)

func MatcherGroup_conn(line <-chan string, result chan<- *Res_Conn, wg_main *sync.WaitGroup) {
	var res_conn = new(Res_Conn)
	res_conn.Init()
	var Matchers_conn = []MatcherType{
		MatcherType{Matcher_automation, &res_conn.automation},
		MatcherType{Matcher_monitoring, &res_conn.monitoring},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go RegexMatchers(Matchers_conn, line, &wg)
	wg.Wait()
	result <- res_conn
	close(result)
	wg_main.Done()
}

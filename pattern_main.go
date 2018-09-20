package main

type Res_Main struct {
	restarts string
}

var res_main = new(Res_Main)

var Matcher_restarts = RegexMatcher_count(`SERVER RESTARTED`)

var Matchers_main = []MatcherType{
	MatcherType{Matcher_restarts, &res_main.restarts},
}

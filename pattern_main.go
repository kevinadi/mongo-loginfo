package main

type Res_Main struct {
	restarts int
}

var res_main = new(Res_Main)

var func_array_main = []RegexMatcher_fn{
	Match_count(`SERVER RESTARTED`, &res_main.restarts),
}

package main

type Res_Conn struct {
	automation string
}

var res_conn = new(Res_Conn)

var Matcher_automation = RegexMatcher_string(`Successfully authenticated as principal (mms-automation)`)

var Matchers_conn = []MatcherType{
	MatcherType{Matcher_automation, &res_conn.automation},
}

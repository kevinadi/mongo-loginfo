package main

type Res_Conn struct {
	automation string
}

var res_conn = new(Res_Conn)

var func_array_conn = []RegexMatcher_fn{
	Match_string(`Successfully authenticated as principal (mms-automation)`, &res_conn.automation),
}

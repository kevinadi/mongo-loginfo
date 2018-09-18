package main

type Res_InitAndListen struct {
	host           string
	port           string
	db_version     string
	storage_engine string
	keyfile        string
	audit          bool
	enterprise     bool
}

var res_initandlisten = new(Res_InitAndListen)

var func_array_initandlisten = []Regex_matcher_fn{
	Match_string(`host=(?P<host>[^\s]+)`, &res_initandlisten.host),
	Match_string(`port=(?P<port>\d+)`, &res_initandlisten.port),
	Match_string(`db version v(?P<dbver>\d{0,2}\.\d{0,2}\.\d{0,2})`, &res_initandlisten.db_version),
	Match_string(`engine:\ *"(?P<storage>[^"]+)"`, &res_initandlisten.storage_engine),
	Match_string(`keyFile:\ *"(?P<keyfile>[^"]+)"`, &res_initandlisten.keyfile),
	Match_bool(`auditLog:\ *{\ *destination:`, &res_initandlisten.audit),
	Match_bool(`modules:\ *enterprise`, &res_initandlisten.enterprise),
}

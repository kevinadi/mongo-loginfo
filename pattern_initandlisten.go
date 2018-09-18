package main

var func_array_initandlisten = []Regex_matcher_fn{
	Match_string(`host=(?P<host>[^\s]+)`, &output.host),
	Match_string(`port=(?P<port>\d+)`, &output.port),
	Match_string(`db version v(?P<dbver>\d{0,2}\.\d{0,2}\.\d{0,2})`, &output.db_version),
	Match_string(`engine:\ *"(?P<storage>[^"]+)"`, &output.storage_engine),
	Match_string(`keyFile:\ *"(?P<keyfile>[^"]+)"`, &output.keyfile),
	Match_bool(`auditLog:\ *{\ *destination:`, &output.audit),
	Match_bool(`modules:\ *enterprise`, &output.enterprise),
}

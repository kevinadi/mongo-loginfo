package main

import "sync"

type Res_InitAndListen struct {
	host           string
	port           string
	db_version     string
	storage_engine string
	keyfile        string
	audit          string
	enterprise     string
	encrypted      string
}

var Matcher_host = RegexMatcher_string(`starting.*host=([^\s]+)`)
var Matcher_port = RegexMatcher_string(`starting.*port=(\d+)`)
var Matcher_storage_engine_1 = RegexMatcher_string(`options:.*storage:.*(wiredTiger|mmapv1)`)
var Matcher_storage_engine_2 = RegexMatcher_string(`(wiredtiger)_open config:`)
var Matcher_db_version = RegexMatcher_string(`db version v(\d{0,2}\.\d{0,2}\.\d{0,2})`)
var Matcher_keyfile = RegexMatcher_string(`options:.*keyFile:\ *"([^"]+)"`)
var Matcher_audit = RegexMatcher_string(`options:.*auditLog:\ *{\ *destination:\ *"([^"]+)"`)
var Matcher_enterprise = RegexMatcher_bool(`modules:\ *enterprise`)
var Matcher_encryption = RegexMatcher_bool(`options:.*enableEncryption:\ *true`)

func MatcherGroup_initandlisten(line <-chan string, result chan<- *Res_InitAndListen, wg_main *sync.WaitGroup) {
	var res_initandlisten = new(Res_InitAndListen)
	var Matchers_initandlisten = []MatcherType{
		MatcherType{Matcher_host, &res_initandlisten.host},
		MatcherType{Matcher_port, &res_initandlisten.port},
		MatcherType{Matcher_storage_engine_1, &res_initandlisten.storage_engine},
		MatcherType{Matcher_storage_engine_2, &res_initandlisten.storage_engine},
		MatcherType{Matcher_db_version, &res_initandlisten.db_version},
		MatcherType{Matcher_keyfile, &res_initandlisten.keyfile},
		MatcherType{Matcher_audit, &res_initandlisten.audit},
		MatcherType{Matcher_enterprise, &res_initandlisten.enterprise},
		MatcherType{Matcher_encryption, &res_initandlisten.encrypted},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go RegexMatchers(Matchers_initandlisten, line, &wg)
	wg.Wait()
	result <- res_initandlisten
	close(result)
	wg_main.Done()
}

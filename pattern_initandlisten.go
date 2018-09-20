package main

import (
	"regexp"
	"sync"
)

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

var res_initandlisten = new(Res_InitAndListen)

type RegexMatcher_func func(result *string, line <-chan string, wg *sync.WaitGroup)
type RegexMatcher_bool_fn func(result *bool, line <-chan string, wg *sync.WaitGroup)

func RegexMatcher_string(regex string, result *string) RegexMatcher_func {
	return func(result *string, line <-chan string, wg *sync.WaitGroup) {
		re := regexp.MustCompile(regex)
		for val := range line {
			matches := re.FindStringSubmatch(val)
			if len(matches) > 0 {
				*result = matches[1]
			}
		}
		wg.Done()
	}
}

func RegexMatcher_bool(regex string, result *string) RegexMatcher_func {
	return func(result *string, line <-chan string, wg *sync.WaitGroup) {
		re := regexp.MustCompile(regex)
		for val := range line {
			matches := re.FindString(val)
			if matches != "" {
				*result = "true"
			}
		}
		wg.Done()
	}
}

var Matcher_host = RegexMatcher_string(`starting.*host=([^\s]+)`, &res_initandlisten.host)
var Matcher_port = RegexMatcher_string(`starting.*port=(\d+)`, &res_initandlisten.port)
var Matcher_storage_engine_1 = RegexMatcher_string(`options:.*storage:.*(wiredTiger|mmapv1)`, &res_initandlisten.storage_engine)
var Matcher_storage_engine_2 = RegexMatcher_string(`(wiredtiger)_open config:`, &res_initandlisten.storage_engine)
var Matcher_db_version = RegexMatcher_string(`db version v(\d{0,2}\.\d{0,2}\.\d{0,2})`, &res_initandlisten.db_version)
var Matcher_keyfile = RegexMatcher_string(`options:.*keyFile:\ *"([^"]+)"`, &res_initandlisten.keyfile)
var Matcher_audit = RegexMatcher_string(`options:.*auditLog:\ *{\ *destination:\ *"([^"]+)"`, &res_initandlisten.audit)
var Matcher_enterprise = RegexMatcher_bool(`modules:\ *enterprise`, &res_initandlisten.enterprise)
var Matcher_encryption = RegexMatcher_bool(`options:.*enableEncryption:\ *true`, &res_initandlisten.encrypted)

var func_array_initandlisten = []RegexMatcher_func{
	Matcher_host,
	Matcher_port,
	Matcher_storage_engine_1,
	Matcher_storage_engine_2,
	Matcher_db_version,
	Matcher_keyfile,
	Matcher_audit,
	Matcher_enterprise,
	Matcher_encryption,
}

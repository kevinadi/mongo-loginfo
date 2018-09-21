package main

import (
	"fmt"
	"sync"
)

type Res_InitAndListen struct {
	host           string
	port           string
	db_version     string
	storage_engine string
	auth           string
	auth_type      string
	keyfile        string
	audit          string
	enterprise     string
	encrypted      string
}

func (r *Res_InitAndListen) String() string {
	return fmt.Sprintf(`
          host: %v
          port: %v
    db_version: %v
storage_engine: %v
          auth: %v
     auth_type: %v
       keyfile: %v
     encrypted: %v
    enterprise: %v
         audit: %v`,
		r.host,
		r.port,
		r.db_version,
		r.storage_engine,
		r.auth,
		r.auth_type,
		r.keyfile,
		r.encrypted,
		r.enterprise,
		r.audit,
	)
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
	var res = new(Res_InitAndListen)
	var Matchers_initandlisten = []MatcherType{
		MatcherType{Matcher_host, &res.host},
		MatcherType{Matcher_port, &res.port},
		MatcherType{Matcher_storage_engine_1, &res.storage_engine},
		MatcherType{Matcher_storage_engine_2, &res.storage_engine},
		MatcherType{Matcher_db_version, &res.db_version},
		MatcherType{Matcher_keyfile, &res.keyfile},
		MatcherType{Matcher_audit, &res.audit},
		MatcherType{Matcher_enterprise, &res.enterprise},
		MatcherType{Matcher_encryption, &res.encrypted},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go RegexMatchers(Matchers_initandlisten, line, &wg)
	wg.Wait()
	result <- res
	close(result)

	if res.keyfile != "" {
		res.auth = "true"
		res.auth_type = "keyfile"
	}
	wg_main.Done()
}

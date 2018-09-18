package main

import (
	"fmt"
	"regexp"
)

var patterns_initandlisten = []GenericPattern{
	GenericPattern{
		"db_version",
		regexp.MustCompile(`db version v(?P<dbver>\d{0,2}\.\d{0,2}\.\d{0,2})`),
		pattern_match,
	},
	GenericPattern{
		"audit",
		regexp.MustCompile(`auditLog:\ *{\ *destination:`),
		pattern_true,
	},
	GenericPattern{
		"using_keyfile",
		regexp.MustCompile(`keyFile:\ *"(?P<keyfile>[^"]+)"`),
		pattern_match,
	},
	GenericPattern{
		"storage_engine",
		regexp.MustCompile(`engine:\ *"(?P<storage>[^"]+)"`),
		pattern_match,
	},
	GenericPattern{
		"enterprise_module",
		regexp.MustCompile(`modules:\ *enterprise`),
		pattern_true,
	},
	GenericPattern{
		"host",
		regexp.MustCompile(`host=(?P<host>[^\s]+)`),
		pattern_match,
	},
	GenericPattern{
		"port",
		regexp.MustCompile(`port=(?P<port>\d+)`),
		pattern_match,
	},
}

func process_initandlisten(line string) {
	for _, pat := range patterns_initandlisten {
		match := pat.regex.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		if pat.title == "host" {
			fmt.Println()
		}
		fmt.Println(pat.title, ":", pat.matchfunc(match))
	}
}

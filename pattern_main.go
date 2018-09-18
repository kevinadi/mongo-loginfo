package main

import (
	"fmt"
	"regexp"
)

var patterns_main = []GenericPattern{
	GenericPattern{
		"restart",
		regexp.MustCompile(`SERVER RESTARTED`),
		pattern_true,
	},
}

func process_main(line string) {
	for _, pat := range patterns_main {
		match := pat.regex.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		fmt.Println(pat.title, ":", pat.matchfunc(match))
	}
}

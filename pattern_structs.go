package main

import (
	"regexp"
)

type GenericPattern struct {
	title     string
	regex     *regexp.Regexp
	matchfunc func([]string) string
}

type ProcessFunc struct {
	cond    string
	process func(string)
}

func pattern_match(match []string) string {
	return match[len(match)-1]
}

func pattern_true(match []string) string {
	return "true"
}

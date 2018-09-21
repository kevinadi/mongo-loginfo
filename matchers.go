package main

import (
	"regexp"
	"strconv"
	"sync"
)

type RegexMatcher_func func(result *string, line <-chan string, wg *sync.WaitGroup)

type MatcherType struct {
	function RegexMatcher_func
	target   *string
}

func RegexMatcher_string(regex string) RegexMatcher_func {
	re := regexp.MustCompile(regex)
	return func(result *string, line <-chan string, wg *sync.WaitGroup) {
		for val := range line {
			matches := re.FindStringSubmatch(val)
			if len(matches) > 0 {
				*result = matches[1]
			}
		}
		wg.Done()
	}
}

func RegexMatcher_bool(regex string) RegexMatcher_func {
	re := regexp.MustCompile(regex)
	return func(result *string, line <-chan string, wg *sync.WaitGroup) {
		for val := range line {
			matches := re.FindString(val)
			if matches != "" {
				*result = "true"
			}
		}
		wg.Done()
	}
}

func RegexMatcher_count(regex string) RegexMatcher_func {
	re := regexp.MustCompile(regex)
	return func(result *string, line <-chan string, wg *sync.WaitGroup) {
		for val := range line {
			matches := re.FindString(val)
			if matches != "" {
				if *result == "" {
					*result = "0"
				}
				num, err := strconv.Atoi(*result)
				if err != nil {
					panic(err)
				}
				*result = strconv.Itoa(num + 1)
			}
		}
		wg.Done()
	}
}

func RegexMatchers(matcher_array []MatcherType, line <-chan string, wg_main *sync.WaitGroup) {
	var wg sync.WaitGroup
	var chans []chan string
	wg.Add(len(matcher_array))

	for i, fn := range matcher_array {
		chans = append(chans, make(chan string, 8))
		go fn.function(fn.target, chans[i], &wg)
	}

	for val := range line {
		for _, chn := range chans {
			chn <- val
		}
	}
	for _, chn := range chans {
		close(chn)
	}

	wg.Wait()
	wg_main.Done()
}

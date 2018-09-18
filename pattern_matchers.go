package main

import (
	"regexp"
	"sync"
)

type Regex_matcher_fn func(<-chan string, *sync.WaitGroup)

func Match_string(regex string, result *string) Regex_matcher_fn {
	return func(line <-chan string, wg *sync.WaitGroup) {
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

func Match_bool(regex string, result *bool) Regex_matcher_fn {
	return func(line <-chan string, wg *sync.WaitGroup) {
		re := regexp.MustCompile(regex)
		for val := range line {
			matches := re.FindString(val)
			if matches != "" {
				*result = true
			}
		}
		wg.Done()
	}
}

func Match_count(regex string, result *int) Regex_matcher_fn {
	return func(line <-chan string, wg *sync.WaitGroup) {
		re := regexp.MustCompile(regex)
		for val := range line {
			matches := re.FindString(val)
			if matches != "" {
				*result += 1
			}
		}
		wg.Done()
	}
}

func Matcher(func_array []Regex_matcher_fn, line <-chan string, output *Output, wg_main *sync.WaitGroup) {
	var wg sync.WaitGroup
	var chans []chan string
	wg.Add(len(func_array))

	for i, fn := range func_array {
		chans = append(chans, make(chan string))
		go fn(chans[i], &wg)
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

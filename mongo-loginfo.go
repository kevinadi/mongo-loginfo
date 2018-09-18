package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var processes = []ProcessFunc{
	ProcessFunc{
		"[main]",
		process_main,
	},
	ProcessFunc{
		"[initandlisten]",
		process_initandlisten,
	},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs a file name")
		os.Exit(1)
	}
	filename := os.Args[1]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var time_start time.Time
	var time_cur string
	var linecount int
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		lineFields := strings.Fields(line)
		if len(lineFields) < 4 {
			continue
		}

		linecount += 1

		if pattern_date.regex.MatchString(lineFields[0]) {
			if time_start.IsZero() {
				time_start = pattern_date.matchfunc(lineFields[0])
			}
			time_cur = lineFields[0]
		}

		for _, proc := range processes {
			if lineFields[3] == proc.cond {
				proc.process(line)
			}
		}

	}

	time_end, _ := time.Parse("2006-01-02T15:04:05-0700", time_cur)
	duration := time_end.Sub(time_start)
	fmt.Println()
	fmt.Println("log start:", time_start)
	fmt.Println("log end:", time_end)
	fmt.Println("log duration:", duration)
	fmt.Println("log length:", linecount)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Output struct {
	filename       string
	db_version     string
	audit          bool
	keyfile        string
	storage_engine string
	enterprise     bool
	host           string
	port           string
	log_start      time.Time
	log_end        time.Time
	log_duration   time.Duration
	log_length     int
	restarts       int
}

func (o *Output) print_output() {
	outstr := `
      filename : %s
          host : %s
          port : %s
    db version : %s
         audit : %t
       keyfile : %s
storage engine : %s
    enterprise : %t

     log start : %v
       log end : %v
  log duration : %v
    log length : %d lines

      restarts : %d
`

	fmt.Printf(outstr,
		o.filename,
		o.host,
		o.port,
		o.db_version,
		o.audit,
		o.keyfile,
		o.storage_engine,
		o.enterprise,
		o.log_start,
		o.log_end,
		o.log_duration,
		o.log_length,
		o.restarts,
	)
}

var output = new(Output)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs a file name")
		os.Exit(1)
	}
	filename := os.Args[1]
	output.filename = filename

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

	ch_ts := make(chan string)
	ch_initandlisten := make(chan string)
	ch_main := make(chan string)

	var wg_main sync.WaitGroup
	var linecount int
	var time_start, time_end time.Time

	go Matcher_timestamp(ch_ts, &time_end, &wg_main)
	go Matcher(func_array_initandlisten, ch_initandlisten, output, &wg_main)
	go Matcher(func_array_main, ch_main, output, &wg_main)
	wg_main.Add(3)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		lineFields := strings.Fields(line)
		if len(lineFields) < 4 {
			continue
		}

		linecount += 1
		if linecount == 1 {
			time_start = parse_timestamp(lineFields[0])
		}
		ch_ts <- lineFields[0]

		switch lineFields[3] {
		case "[initandlisten]":
			ch_initandlisten <- strings.Join(lineFields[4:], " ")
		case "[main]":
			ch_main <- strings.Join(lineFields[4:], " ")
		}

	}

	close(ch_ts)
	close(ch_initandlisten)
	close(ch_main)
	wg_main.Wait()

	output.log_start = time_start
	output.log_end = time_end
	output.log_duration = time_end.Sub(time_start)
	output.log_length = linecount

	output.print_output()

}

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

func Read_file(filename string, line chan<- string) {
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
	for scanner.Scan() {
		line <- scanner.Text()
	}
	close(line)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs a file name")
		os.Exit(1)
	}
	filename := os.Args[1]
	output.filename = filename

	var wg_main sync.WaitGroup
	var linecount int
	var time_start, time_end time.Time

	ch_line := make(chan string)
	chans := map[string]chan string{
		"ts":            make(chan string),
		"initandlisten": make(chan string),
		"main":          make(chan string),
	}

	go Read_file(filename, ch_line)
	go Matcher_timestamp(chans["ts"], &time_end, &wg_main)
	go Matcher(func_array_initandlisten, chans["initandlisten"], output, &wg_main)
	go Matcher(func_array_main, chans["main"], output, &wg_main)
	wg_main.Add(len(chans))

	for line := range ch_line {

		lineFields := strings.Fields(line)
		if len(lineFields) < 4 {
			continue
		}

		linecount += 1
		if linecount == 1 {
			time_start = parse_timestamp(lineFields[0])
		}
		chans["ts"] <- lineFields[0]

		switch lineFields[3] {
		case "[initandlisten]":
			chans["initandlisten"] <- strings.Join(lineFields[4:], " ")
		case "[main]":
			chans["main"] <- strings.Join(lineFields[4:], " ")
		}

	}

	for _, ch := range chans {
		close(ch)
	}
	wg_main.Wait()

	output.log_start = time_start
	output.log_end = time_end
	output.log_duration = time_end.Sub(time_start)
	output.log_length = linecount

	output.print_output()

}

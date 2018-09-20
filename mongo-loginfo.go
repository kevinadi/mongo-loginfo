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
	filename      string
	log_start     time.Time
	log_end       time.Time
	log_duration  time.Duration
	log_length    int
	initandlisten *Res_InitAndListen
	main          *Res_Main
	conn          *Res_Conn
}

func (o *Output) print_output() {
	outstr := `
      filename : %v
          host : %v
          port : %v
     log start : %v
       log end : %v
  log duration : %v
    log length : %v lines
    db version : %v
storage engine : %v

Features
  Audit      : %v
  Keyfile    : %v
  Enterprise : %v
  Automation : %v
  Encryption : %v

Events
  Restarts   : %v
`
	fmt.Printf(outstr,
		o.filename,
		o.initandlisten.host,
		o.initandlisten.port,
		o.log_start,
		o.log_end,
		o.log_duration,
		o.log_length,
		o.initandlisten.db_version,
		o.initandlisten.storage_engine,
		o.initandlisten.audit,
		o.initandlisten.keyfile,
		o.initandlisten.enterprise,
		o.conn.automation,
		o.initandlisten.encrypted,
		o.main.restarts,
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
		"conn":          make(chan string),
	}

	go Read_file(filename, ch_line)
	go Matcher_timestamp(chans["ts"], &time_end, &wg_main)
	//go Matcher(func_array_initandlisten, chans["initandlisten"], output, &wg_main)
	go Matcher(func_array_main, chans["main"], output, &wg_main)
	go Matcher(func_array_conn, chans["conn"], output, &wg_main)

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

		switch {
		case lineFields[3] == "[initandlisten]":
			chans["initandlisten"] <- strings.Join(lineFields[4:], " ")
		case lineFields[3] == "[main]":
			chans["main"] <- strings.Join(lineFields[4:], " ")
		case strings.HasPrefix(lineFields[3], "[conn"):
			chans["conn"] <- strings.Join(lineFields[4:], " ")
		}

	}

	for _, ch := range chans {
		close(ch)
	}
	wg_main.Wait()

	output.initandlisten = res_initandlisten
	output.main = res_main
	output.conn = res_conn
	output.log_start = time_start
	output.log_end = time_end
	output.log_duration = time_end.Sub(time_start)
	output.log_length = linecount

	output.print_output()

}

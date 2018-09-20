package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Output struct {
	filename      string
	initandlisten *Res_InitAndListen
	main          *Res_Main
	conn          *Res_Conn
	logtimes      *Res_LogTimes
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
		o.logtimes.log_start,
		o.logtimes.log_end,
		o.logtimes.log_duration,
		o.logtimes.log_length,
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
	var GlobalOutput = new(Output)

	if len(os.Args) < 2 {
		fmt.Println("Needs a file name")
		os.Exit(1)
	}
	filename := os.Args[1]
	GlobalOutput.filename = filename

	var wg_main sync.WaitGroup

	ch_line := make(chan string)
	chans := map[string]chan string{
		"ts":            make(chan string, 8),
		"initandlisten": make(chan string, 8),
		"main":          make(chan string, 8),
		"conn":          make(chan string, 8),
	}

	go Read_file(filename, ch_line)

	output_logtimes := make(chan *Res_LogTimes)
	go Matcher_timestamp(chans["ts"], output_logtimes, &wg_main)

	output_initandlisten := make(chan *Res_InitAndListen)
	go MatcherGroup_initandlisten(chans["initandlisten"], output_initandlisten, &wg_main)

	output_main := make(chan *Res_Main)
	go MatcherGroup_main(chans["main"], output_main, &wg_main)

	output_conn := make(chan *Res_Conn)
	go MatcherGroup_conn(chans["conn"], output_conn, &wg_main)

	wg_main.Add(len(chans))

	for line := range ch_line {

		lineFields := strings.Fields(line)
		if len(lineFields) < 4 {
			continue
		}

		chans["ts"] <- lineFields[0]

		restofline := strings.Join(lineFields[4:], " ")
		switch {
		case lineFields[3] == "[initandlisten]":
			chans["initandlisten"] <- restofline
		case lineFields[3] == "[main]":
			chans["main"] <- restofline
		case strings.HasPrefix(lineFields[3], "[conn"):
			chans["conn"] <- restofline
		}

	}

	for _, ch := range chans {
		close(ch)
	}
	GlobalOutput.logtimes = <-output_logtimes
	GlobalOutput.main = <-output_main
	GlobalOutput.conn = <-output_conn
	GlobalOutput.initandlisten = <-output_initandlisten
	wg_main.Wait()

	GlobalOutput.print_output()
}

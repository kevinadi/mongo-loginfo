package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
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

func (o *Output) String() string {
	return fmt.Sprintf(`mongo-loginfo %v %v
 
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
   Authorization : %v
  Authentication : %v
         Keyfile : %v
           Audit : %v
      Enterprise : %v
      Automation : %v
      Monitoring : %v
      Encryption : %v
 
Events
        Restarts : %v
        fasserts : %v`,
		version,
		date,
		o.filename,
		o.initandlisten.host,
		o.initandlisten.port,
		o.logtimes.log_start,
		o.logtimes.log_end,
		o.logtimes.log_duration,
		o.logtimes.log_length,
		o.initandlisten.db_version,
		o.initandlisten.storage_engine,

		o.initandlisten.auth,
		o.initandlisten.auth_type,
		o.initandlisten.keyfile,
		o.initandlisten.audit,
		o.initandlisten.enterprise,
		o.conn.automation,
		o.conn.monitoring,
		o.initandlisten.encrypted,
		o.main.restarts,
		o.main.fasserts,
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

var (
	version string
	date    string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	flag.Usage = func() {
		fmt.Printf("Usage: %s [FILE] [-version] [-help]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Bool("help", false, "Show help information")
	versionPtr := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *versionPtr {
		fmt.Println(os.Args[0], version, date)
		os.Exit(0)
	}
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var wg_main sync.WaitGroup
	var GlobalOutput = new(Output)

	filename := flag.Args()[0]
	GlobalOutput.filename = filename

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

		switch {
		case lineFields[3] == "[initandlisten]":
			chans["initandlisten"] <- line
		case strings.HasPrefix(lineFields[3], "[conn"):
			switch lineFields[2] {
			case "ACCESS":
				chans["conn"] <- line
			case "CONTROL":
				chans["initandlisten"] <- line
			}
		default:
			chans["main"] <- line
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

	fmt.Println(GlobalOutput)
}

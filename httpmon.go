package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	intervalS = 10
	maxbuf    = 50
)

func main() {
	var stdin bool
	flag.BoolVar(&stdin, "stdin", false, "Use stdin")
	var fin string
	flag.StringVar(&fin, "file", "", "Use file input")
	var tthresh int
	flag.IntVar(&tthresh, "threshold", 10, "Default traffic threshold")
	flag.Parse()

	if !stdin && fin == "" {
		log.Fatalf("need 1-2 args, i for interactive with stdin, f {filepath} for file input\n")
		os.Exit(1)
	}

	// var ir io.Reader
	if stdin {
		log.Printf("Interactive mode, input on stdin\n")
		ir := bufio.NewReader(os.Stdin)
		monitor(ir)
	} else {
		log.Printf("File mode, reading from %s\n", fin)
		ir, err := os.Open(fin)
		if err != nil {
			log.Fatal(err)
		}
		defer ir.Close()
		monitor(ir)
	}

}

func monitor(ir io.Reader) int {
	// # 10s segments = 2 mins * (6 segments in 1 min)
	m := NewMonitor(6 * 2)
	log.Printf("mon %v", m)

	s := bufio.NewScanner(ir)
	log.Printf("S %v", s)

	var readHeader sync.Once
	readHeader.Do(func() {
		if s.Scan() {
			log.Printf("Header: %v", s.Text())
		} else {
			log.Fatalln("Somehow failed to read header first line.")
		}
	})
	for s.Scan() {
		log.Printf("Processing: %v", s.Text())
		linfo := strings.Split(s.Text(), ",")
		// timestamp
		lts, err := strconv.ParseInt(linfo[3], 10, 64)
		if err != nil {
			log.Printf("Faile to get timestamp, dropping %v", linfo)
			continue
		}
		m.ProcessLine(lts)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}

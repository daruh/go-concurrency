package main

import (
	"fmt"
	"go-concurrency/powertools/safemap"
	"os"
	"path/filepath"
	"runtime"
)

//go run .\apachereport.go ..\logsData\access_test.log

var workers = runtime.NumCPU()

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file.log>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	lines := make(chan string, workers*4)
	done := make(chan struct{}, workers)
	pageMap := safemap.New()

	go readLines(os.Args[1], lines)
	processLines(done, pageMap, lines)
	waitUntil(done)
	showResults(pageMap)

}

func readLines(filename string, lines chan<- string) {

}

func processLines(done chan<- struct{}, pageMap safemap.SafeMap, lines <-chan string) {

}

func waitUntil(done <-chan struct{}) {

}

func showResults(pageMap safemap.SafeMap) {

}

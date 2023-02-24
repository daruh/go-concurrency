package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
)

var workers = runtime.NumCPU()

//go run .\apachereport.go ../logsData/access_test.log

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file.log>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	lines := make(chan string, workers*4)
	done := make(chan struct{}, workers)
	pageMapObj := NewPageMap()
	go readLines(os.Args[1], lines)
	getRx := regexp.MustCompile(`GET[ \t]+([^ \t\n]+[.]php?)`)
	for i := 0; i < workers; i++ {
		go processLines(done, getRx, pageMapObj, lines)
	}
	waitUntil(done)
	showResults(pageMapObj)
}

func readLines(filename string, lines chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open the file:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines <- line
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file:", err)
			}
			break
		}
	}
	close(lines)
}

func processLines(done chan<- struct{}, getRx *regexp.Regexp, pageMap *pageMap, lines <-chan string) {
	for line := range lines {
		if matches := getRx.FindStringSubmatch(line); matches != nil {
			pageMap.Increment(matches[1])
		}
	}
	done <- struct{}{}
}
func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}

func showResults(pageMap *pageMap) {
	for page, count := range pageMap.counterForPage {
		fmt.Printf("%8d %s\n", count, page)
	}
}

type pageMap struct {
	counterForPage map[string]int
	mutex          *sync.RWMutex
}

func NewPageMap() *pageMap {
	return &pageMap{make(map[string]int), new(sync.RWMutex)}
}

func (pm *pageMap) Increment(page string) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.counterForPage[page]++
}

func (pm pageMap) Len() int {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	return len(pm.counterForPage)
}

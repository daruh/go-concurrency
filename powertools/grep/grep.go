package main

import (
	"fmt"
	"go-concurrency/powertools/utils"
	"regexp"
	"runtime"
)

type Result struct {
	filename string
	lino     int
	line     string
}

var (
	workers = runtime.NumCPU()
)

func grep(lineRx *regexp.Regexp, filenames []string) {

	jobs := make(chan Job, workers)
	results := make(chan Result, utils.Minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results) //blocks
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs) //blocks
	}
	go awaitCompletion(done, results)
	processResults(results)
}

func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename: filename, results: results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{}
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
	}
}

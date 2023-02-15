package main

import (
	"fmt"
	"go-concurrency/powertools/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		var timeout = time.Second
		grep(timeout, lineRx, utils.CommandLineFiles(os.Args[2:]))
	}
}

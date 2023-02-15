package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Filter tool...")
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	log.SetFlags(0)

	//-algorithm 1 -suffixes ".pdf" "file1.txt" "report.pdf" "document.xlsx"
	algorithm, minSize, maxSize, suffixes, files := handleCommandLine()

	if algorithm == 1 {
		Sink(FilterSize(minSize, maxSize, FilterSuffixes(suffixes, Source(files))))
	} else {
		channel1 := Source(files)
		channel2 := FilterSuffixes(suffixes, channel1)
		channel3 := FilterSize(minSize, maxSize, channel2)
		Sink(channel3)
	}
}

func handleCommandLine() (algorithm int, minSize, maxSize int64, suffixes, files []string) {
	flag.IntVar(&algorithm, "algorithm", 1, "1 or 2")
	flag.Int64Var(&minSize, "min", -1, "minimum file size (-1 means no minimum)")
	flag.Int64Var(&maxSize, "max", -1, "maximum file size (-1 means no minimum)")
	var suffixesOpt = flag.String("suffixes", "", "comma-separated list of file suffixes")
	flag.Parse()
	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be < maximum size")
	}
	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}
	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

func Source(files []string) <-chan string {

	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()

	return out
}

func FilterSize(minimum, maximum int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))

	go func() {

		for filename := range in {

			if minimum == -1 && maximum == -1 {
				out <- filename
				continue
			}
			finfo, err := os.Stat(filename)
			if err != nil {
				continue
				//ignore files we can't process
			}
			size := finfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) &&
				(maximum == -1 || maximum > -1 && maximum >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func FilterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))

	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}

			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func Sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

package utils

import (
	"os"
	"path/filepath"
	"strings"
)

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

func SinkCollection(in <-chan string) []string {
	list := []string{}
	for filename := range in {
		list = append(list, filename)
	}
	return list
}

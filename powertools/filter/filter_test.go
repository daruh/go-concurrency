package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func SinkCollection(in <-chan string) []string {
	list := []string{}
	for filename := range in {
		list = append(list, filename)
	}
	return list
}

func TestFilter(t *testing.T) {
	files := []string{"file1.txt", "report.pdf", "document.xlsx"}

	suffixes := []string{".pdf"}

	list := SinkCollection(FilterSize(-1, -1, FilterSuffixes(suffixes, Source(files))))
	assert.Equal(t, 1, len(list))
	assert.Equal(t, "report.pdf", list[0])
}

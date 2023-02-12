package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	files := []string{"file1.txt", "report.pdf", "document.xlsx"}

	suffixes := []string{".pdf"}

	filterFiles := sink(filterSuffixes(suffixes, source(files)))

	assert.Equal(t, 3, filterFiles)
}

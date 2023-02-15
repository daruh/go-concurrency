package main

import (
	"go-concurrency/powertools/utils"
	"regexp"
	"testing"
)

func TestGrep(t *testing.T) {

	filenames, _ := utils.GetFilesFromDirectory("../testData", []string{".txt"})
	regex, _ := regexp.Compile("darek")
	grep(regex, filenames)
}

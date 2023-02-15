package main

import (
	"go-concurrency/powertools/utils"
	"regexp"
	"testing"
	"time"
)

func TestGrep_Success(t *testing.T) {

	filenames, _ := utils.GetFilesFromDirectory("../testData", []string{".txt"})
	regex, _ := regexp.Compile("darek")
	grep(time.Minute*10, regex, filenames)
}

func TestGrep_Timeout(t *testing.T) {

	filenames, _ := utils.GetFilesFromDirectory("../testData", []string{".txt"})
	regex, _ := regexp.Compile("darek")
	grep(time.Millisecond, regex, filenames)
}

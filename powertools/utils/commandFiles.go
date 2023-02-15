package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func CommandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func GetFilesFromDirectory(directory string, suffixes []string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, directory+"\\"+file.Name())
	}

	return SinkCollection(FilterSize(-1, -1, FilterSuffixes(suffixes, Source(filenames)))), nil
}

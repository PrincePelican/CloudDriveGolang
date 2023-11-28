package validator

import (
	"mime/multipart"
	"path/filepath"
	"regexp"
)

func isFilePathValid(path string) bool {
	r1 := regexp.MustCompile(`^/`)
	r2 := regexp.MustCompile(`/$`)
	r3 := regexp.MustCompile(`/\.\./`)
	r4 := regexp.MustCompile(`\\`)
	return !r1.MatchString(path) && !r2.MatchString(path) && !r3.MatchString(path) && !r4.MatchString(path)
}

func isFilesMatchingPaths(files []*multipart.FileHeader, paths []string) bool {
	for index, path := range paths {
		fileName := filepath.Base(path)
		if files[index].Filename != fileName {
			return false
		}
	}
	return true
}

func ValidateFileCreateForm(files []*multipart.FileHeader, paths []string) bool {
	if len(files) != len(paths) {
		return false
	}
	for _, filePath := range paths {
		if !isFilePathValid(filePath) {
			return false
		}
	}

	return isFilesMatchingPaths(files, paths)
}

package validator

import (
	"mime/multipart"
	"testing"
)

func TestIsFilePathValidWhenFullPath(t *testing.T) {
	expected := true
	path := "good/valid/test.txt"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be valid", path)
	}
}

func TestIsFilePathValidWhenSingleFile(t *testing.T) {
	expected := true
	path := "test.txt"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be valid", path)
	}
}

func TestIsFilePathInvalid(t *testing.T) {
	expected := false
	path := "/bad/text.txt"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be invalid", path)
	}
}

func TestIsFilePathInvalidWithoutFileExtention(t *testing.T) {
	expected := false
	path := "/path"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be invalid", path)
	}
}

func TestIsFilePathInvalidWithDots(t *testing.T) {
	expected := false
	path := "/../path.txt"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be invalid", path)
	}
}

func TestIsFilePathInvalidWithOtherSlashes(t *testing.T) {
	expected := false
	path := "\\path.txt"
	actual := isFilePathValid(path)
	if actual != expected {
		t.Errorf("path:%s should be invalid", path)
	}
}

func TestSingleCorrectValidateFileCreateForm(t *testing.T) {
	expected := true
	path := "dir/test/test.txt"
	fileName := "test.txt"
	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     4,
	}
	paths := []string{path}
	files := []*multipart.FileHeader{fileHeader}
	actual := ValidateFileCreateForm(files, paths)
	if actual != expected {
		t.Errorf("fileNames:%s paths:%s should be valid", fileName, path)
	}
}

func TestDoublePathsOneFileNameValidateFileCreateForm(t *testing.T) {
	expected := false
	path1 := "dir/test/test.txt"
	path2 := "dir/test/test.txt"
	fileName := "test.txt"
	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     4,
	}
	paths := []string{path1, path2}
	files := []*multipart.FileHeader{fileHeader}
	actual := ValidateFileCreateForm(files, paths)
	if actual != expected {
		t.Errorf("fileNames:%s paths:%s %s should be invalid", fileName, path1, path2)
	}
}

func TestOnePathsDoubleFileNameValidateFileCreateForm(t *testing.T) {
	expected := false
	path := "dir/test/test1.txt"
	fileName1 := "test1.txt"
	fileName2 := "test2.txt"
	fileHeader1 := &multipart.FileHeader{
		Filename: fileName1,
		Size:     4,
	}
	fileHeader2 := &multipart.FileHeader{
		Filename: fileName2,
		Size:     4,
	}
	paths := []string{path}
	files := []*multipart.FileHeader{fileHeader1, fileHeader2}
	actual := ValidateFileCreateForm(files, paths)
	if actual != expected {
		t.Errorf("fileNames:%s %s paths:%s should be invalid", fileName1, fileName2, path)
	}
}

func TestInvalidSingleValidateFileCreateForm(t *testing.T) {
	expected := false
	path := "dir/test/test.txt"
	fileName := "testBad.txt"
	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     4,
	}
	paths := []string{path}
	files := []*multipart.FileHeader{fileHeader}
	actual := ValidateFileCreateForm(files, paths)
	if actual != expected {
		t.Errorf("fileNames:%s paths:%s should be invalid", fileName, path)
	}
}

func TestValidDoublePathsDoubleFileNameValidateFileCreateForm(t *testing.T) {
	expected := true
	path1 := "dir/test/test1.txt"
	path2 := "dir/test/test2.txt"
	fileName1 := "test1.txt"
	fileName2 := "test2.txt"
	fileHeader1 := &multipart.FileHeader{
		Filename: fileName1,
		Size:     4,
	}
	fileHeader2 := &multipart.FileHeader{
		Filename: fileName2,
		Size:     4,
	}
	paths := []string{path1, path2}
	files := []*multipart.FileHeader{fileHeader1, fileHeader2}
	actual := ValidateFileCreateForm(files, paths)
	if actual != expected {
		t.Errorf("fileNames:%s %s paths:%s %s should be valid", fileName1, fileName2, path1, path2)
	}
}

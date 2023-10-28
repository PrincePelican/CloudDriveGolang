package converter

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func getTmp() string {
	var tmp string
	if env := os.Getenv("TmpFolder"); env != "" {
		tmp = env
	}
	return tmp
}

func ZipDir(sourceDir string, zipFilePath string) (string, error) {
	zipFilePath = filepath.Join(getTmp(), zipFilePath+".zip")
	sourceDir = filepath.Join(getTmp(), sourceDir)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.Join(filepath.Base(sourceDir), path[len(sourceDir):])

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		w, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(w, file)
			return err
		}

		return nil
	})

	return zipFilePath, err
}

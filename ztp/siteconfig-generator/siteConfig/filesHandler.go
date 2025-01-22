package siteConfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const fileNameLength = 30

type DirContainFiles struct {
	Directory string
	Files     []os.FileInfo
}

func resolveFilePath(filePath string, basedir string) string {
	if _, errAbsPath := os.Stat(filePath); errAbsPath == nil {
		return filePath
	}
	return basedir + "/" + filePath
}

func GetFiles(path string) ([]os.FileInfo, error) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		var files []os.FileInfo
		results, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		// Translate []fs.DirEntry to []os.FileInfo
		for _, result := range results {
			resultsInfo, err := result.Info()
			if err != nil {
				return nil, err
			}
			files = append(files, resultsInfo)
		}

		return files, nil
	}

	return []os.FileInfo{fileInfo}, nil
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFile(filePath string, outDir string, content []byte) error {
	path := outDir + "/" + filePath[:strings.LastIndex(filePath, "/")]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0775)
	}
	err := os.WriteFile(outDir+"/"+filePath, content, 0644)

	return err
}

func ReadExtraManifestResourceFile(filePath string) ([]byte, error) {
	var dir = ""
	var err error = nil
	var ret []byte

	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	dir = filepath.Dir(ex)
	err = CheckFileName(filePath)
	if err != nil {
		return ret, err
	}
	ret, err = ReadFile(resolveFilePath(filePath, dir))

	// added fail safe for test runs as `os.Executable()` will fail for tests
	if err != nil {

		dir, err = os.Getwd()
		if err != nil {
			return nil, err
		}

		ret, err = ReadFile(resolveFilePath(filePath, dir))
	}
	return ret, err
}

func GetExtraManifestResourceDir(manifestsPath string) (string, error) {

	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(ex)

	return resolveFilePath(manifestsPath, dir), err
}

func GetExtraManifestResourceFiles(manifestsPath string) ([]os.FileInfo, error) {

	var files []os.FileInfo

	dirPath, err := GetExtraManifestResourceDir(manifestsPath)
	if err != nil {
		return files, err
	}

	files, err = GetFiles(dirPath)
	if err != nil {
		dir, err := os.Getwd()

		if err != nil {
			return nil, err
		}

		files, err = GetFiles(resolveFilePath(manifestsPath, dir))
	}
	return files, err
}

func CheckFileName(filePath string) error {
	// Get filename from the filepath
	var fileName = filePath[strings.LastIndex(filePath, "/")+1:]
	// Checking filename length without separator '.'
	fileNameStripped := strings.ReplaceAll(fileName, ".", "")
	if len(fileNameStripped) > fileNameLength {
		return fmt.Errorf("\nfilename too long: %d - %s expected length without separator '.' < %d", len(fileNameStripped), fileName, fileNameLength)
	}
	return nil
}

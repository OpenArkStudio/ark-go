package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath) // For read access.
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

// PathExists : if the file/dir exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Write : write file
func Write(path string, writeBytes []byte) error {
	if PathExists(path) {
		return fmt.Errorf("path[%s] is already exists", path)
	}

	return ForceWrite(path, writeBytes)
}

// ForceWrite overwrite file if exist
func ForceWrite(path string, writeBytes []byte) error {
	return ioutil.WriteFile(path, writeBytes, 0777)
}

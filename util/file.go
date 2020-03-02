package util

import (
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

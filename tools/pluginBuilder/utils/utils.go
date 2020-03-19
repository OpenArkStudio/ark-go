package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"unicode"
)

// utils
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func ParseTemplate(str string, data interface{}) (string, error) {
	t, err := template.New("aaaaaa").Parse(str)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

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

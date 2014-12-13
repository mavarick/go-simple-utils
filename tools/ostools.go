package tools

import (
	"errors"
	"fmt"
	"os"
)

// check if path exists
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// check if path is directory
func IsPathDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

// check if path one file
func IsPathFile(path string) bool {
	return !IsPathDir(path)
}

// Create Directory
func CreateDir(path string) {
	if IsPathExist(path) == true {
		if IsPathDir(path) == true {
			fmt.Println(fmt.Sprintf("Directory [%s] exists", path))
		} else {
			panic(errors.New(fmt.Sprintf("File [%s] exists! ", path)))
		}
	} else {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

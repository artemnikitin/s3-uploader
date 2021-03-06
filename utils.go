package main

import (
	"bytes"
	"strings"
)

func createKey(path string) string {
	var buffer bytes.Buffer
	buffer.WriteString(*uploadPath)
	if *uploadPath == "/" {
		if startWith(path, "/") {
			return path
		}
		buffer.WriteString(path)
	} else {
		if !endWith(*uploadPath, "/") && !startWith(path, "/") {
			buffer.WriteString("/")
		}
		if endWith(*uploadPath, "/") && startWith(path, "/") {
			buffer.WriteString(string(path[1:]))
		} else {
			buffer.WriteString(path)
		}
	}
	return buffer.String()
}

func startWith(original, substring string) bool {
	if len(substring) > len(original) {
		return false
	}
	str := string(original[0:len(substring)])
	return str == substring
}

func endWith(original, substring string) bool {
	if len(substring) > len(original) {
		return false
	}
	str := string(original[len(original)-len(substring):])
	return str == substring
}

func getFileName(filepath string) string {
	if *rename != "" {
		return *rename
	}
	index := strings.LastIndex(filepath, "/")
	if index == -1 {
		return filepath
	}
	return filepath[index+1:]
}

func getFolderName(filepath string) string {
	var result string
	if endWith(filepath, "/") {
		pos := strings.LastIndex(string(filepath[:len(filepath)-1]), "/")
		result = filepath[pos+1 : len(filepath)-1]
	} else {
		pos := strings.LastIndex(filepath, "/")
		result = filepath[pos+1:]
	}
	return result
}

func getPathInsideFolder(path, folder string) string {
	pos := strings.Index(path, folder)
	var result string
	if pos != -1 {
		result = string(path[pos-1:])
	}
	return result
}

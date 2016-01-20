package main

import (
	"log"
	"testing"
)

func TestEndsWithPositive(t *testing.T) {
	str := "dfd/"
	if !endWith(str, "/") {
		t.Error("String ends with / should be true")
	}
}

func TestEndsWithBig(t *testing.T) {
	str := "dfd/"
	if endWith(str, "sdssd/") {
		t.Error("String ends with sdssd/ should be false")
	}
}

func TestEndsWithPositive2Symbols(t *testing.T) {
	str := "dfd/"
	if !endWith(str, "d/") {
		t.Error("String ends with d/ should be true")
	}
}

func TestEndsWithNegative(t *testing.T) {
	str := "erwerwe"
	if endWith(str, "/") {
		t.Error("String ends with / should be false")
	}
}

func TestEndsWithNegative2Symbols(t *testing.T) {
	str := "erwerwe"
	if endWith(str, "d/") {
		t.Error("String ends with d/ should be false")
	}
}

func TestStartWithPositive(t *testing.T) {
	str := "qwe"
	if !startWith(str, "q") {
		t.Error("String should start with q")
	}
}

func TestStartWithBig(t *testing.T) {
	str := "qwe"
	if startWith(str, "qwer") {
		t.Error("String shouldn't start with qwer")
	}
}

func TestStartWithNegative(t *testing.T) {
	str := "abc"
	if startWith(str, "w") {
		t.Error("String shouldn't start with w")
	}
}

func TestStartWithPositive2Symbols(t *testing.T) {
	str := "qwe"
	if !startWith(str, "qw") {
		t.Error("String should start with q")
	}
}

func TestStartWithNegative2Symbols(t *testing.T) {
	str := "abc"
	if startWith(str, "sw") {
		t.Error("String shouldn't start with w")
	}
}

func TestGetFolderName(t *testing.T) {
	path := "/ff/ddd/fff"
	path2 := "/ff/ddd/fff/"
	if getFolderName(path) != "fff" {
		t.Error("Folder name should be fff")
	}
	if getFolderName(path2) != "fff" {
		t.Error("Folder name should be fff")
	}
}

func TestGetPathInsideFolder(t *testing.T) {
	path := "/fff/ddd/dir/sc/23"
	path2 := "/fff/ddd/dir/sc/23/"
	directory := "dir"
	if getPathInsideFolder(path, directory) != "/dir/sc/23" {
		t.Error("Path should be /dir/sc/23")
	}
	if getPathInsideFolder(path2, directory) != "/dir/sc/23/" {
		t.Error("Path should be /dir/sc/23/")
	}
	if getPathInsideFolder(path, "qwerty") != "" {
		t.Error("Path should be blank")
	}
}

func TestCreateKey(t *testing.T) {
	path := "/ddd/qwe/asd.jjj"
	path2 := "ddd/qwe/asd.jjj"

	if createKey(path) != "/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", path)
	}
	if createKey(path2) != "/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", path)
	}

	*uploadpath = "/vvv/ddd/"
	if createKey(path) != "/vvv/ddd/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", *uploadpath+path2)
	}
	if createKey(path2) != "/vvv/ddd/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", *uploadpath+path2)
	}

	*uploadpath = "/vvv/ddd"
	if createKey(path) != "/vvv/ddd/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", *uploadpath+path)
	}
	if createKey(path2) != "/vvv/ddd/ddd/qwe/asd.jjj" {
		t.Error("Created key should be equal to", *uploadpath+path)
	}
}

func TestGetFileName(t *testing.T) {
	*rename = ""

	path := getFileName("dfd.kkk")
	if path != "dfd.kkk" {
		log.Println("path =", path)
		t.Error("Should return dfd.kkk for file")
	}

	path = getFileName("/dfdfd/www/dfd.kkk")
	if path != "dfd.kkk" {
		log.Println("path =", path)
		t.Error("Should return dfd.kkk for path")
	}

	*rename = "dfg.lll"
	path = getFileName("/dfdfd/www/dfd.kkk")
	if path != "dfg.lll" {
		log.Println("path =", path)
		t.Error("Should return dfg.lll if rename option is specified")
	}
}

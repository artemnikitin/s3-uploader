package main

import "testing"

func TestEndWith(t *testing.T) {
	cases := []struct {
		src, sub string
		res      bool
	}{
		{"dfd/", "/", true},
		{"dfd/", "d/", true},
		{"dfd/", "sdssd/", false},
		{"dfd/", "f", false},
		{"dfd", "/", false},
		{"dfd/", "x/", false},
		{"", "d/", false},
		{"dfd/", "", true},
		{"", "", true},
	}

	for _, v := range cases {
		result := endWith(v.src, v.sub)
		if result != v.res {
			t.Errorf("For string: %s end with: %s, actual: %v, expected: %v", v.src, v.sub, result, v.res)
		}
	}
}

func TestStartWith(t *testing.T) {
	cases := []struct {
		src, sub string
		res      bool
	}{
		{"qwe", "q", true},
		{"qwe", "qw", true},
		{"qwe", "qwer", false},
		{"abc", "w", false},
		{"", "d/", false},
		{"dfd/", "", true},
		{"", "", true},
	}

	for _, v := range cases {
		result := startWith(v.src, v.sub)
		if result != v.res {
			t.Errorf("For string: %s start with: %s, actual: %v, expected: %v", v.src, v.sub, result, v.res)
		}
	}
}

func TestGetFolderName(t *testing.T) {
	cases := []struct{ path, folder string }{
		{"/ff/ddd/fff", "fff"},
		{"/ff/ddd/fff/", "fff"},
	}

	for _, v := range cases {
		result := getFolderName(v.path)
		if result != v.folder {
			t.Errorf("For path: %s folder actual: %s, expected: %s", v.path, result, v.folder)
		}
	}
}

func TestGetPathInsideFolder(t *testing.T) {
	cases := []struct{ path, dir, res string }{
		{"/fff/ddd/dir/sc/23", "dir", "/dir/sc/23"},
		{"/fff/ddd/dir/sc/23/", "dir", "/dir/sc/23/"},
		{"/fff/ddd/dir/sc/23/", "qwerty", ""},
	}

	for _, v := range cases {
		result := getPathInsideFolder(v.path, v.dir)
		if result != v.res {
			t.Errorf("For path: %s and directory: %s actual output: %s, expected: %s", v.path, v.dir, result, v.res)
		}
	}
}

func TestCreateKey(t *testing.T) {
	cases := []struct{ uploadPath, path, res string }{
		{"", "/ddd/qwe/asd.jjj", "/ddd/qwe/asd.jjj"},
		{"", "ddd/qwe/asd.jjj", "/ddd/qwe/asd.jjj"},
		{"/vvv/ddd/", "/ddd/qwe/asd.jjj", "/vvv/ddd/ddd/qwe/asd.jjj"},
		{"/vvv/ddd/", "ddd/qwe/asd.jjj", "/vvv/ddd/ddd/qwe/asd.jjj"},
		{"/vvv/ddd", "ddd/qwe/asd.jjj", "/vvv/ddd/ddd/qwe/asd.jjj"},
	}

	for _, v := range cases {
		*uploadpath = v.uploadPath
		result := createKey(v.path)
		if result != v.res {
			t.Errorf("For path: %s actual key: %s, expected: %s", v.path, result, v.res)
		}
	}
}

func TestGetFileName(t *testing.T) {
	cases := []struct{ rename, path, res string }{
		{"", "dfd.kkk", "dfd.kkk"},
		{"", "/dfdfd/www/dfd.kkk", "dfd.kkk"},
		{"dfg.lll", "/dfdfd/www/dft.kkk", "dfg.lll"},
	}

	for _, v := range cases {
		*rename = v.rename
		result := getFileName(v.path)
		if result != v.res {
			t.Errorf("For path: %s actual filename: %s, expected: %s", v.path, result, v.res)
		}
	}
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
	"path/filepath"
)

var abspath func(string) (string, error) = filepath.Abs
var relpath func(string, string) (string, error) = filepath.Rel
var getpath func(s string) string = getAbsPath

func getAbsPath(s string) string {
	var err error
	p := filepath.FromSlash(s)
	p, err = abspath(p)
	if err != nil {
		Panic(err)
	}
	return p
}

func Rmtree(path string) {
	path = getpath(path)
	if st, err := os.Stat(path); err == nil {
		if st.IsDir() {
			Log("rmtree: %q", path)
		} else {
			Panicf("rmtree: %q is not a directory", path)
		}
	}
	if err := fs.RemoveAll(path); err != nil {
		Panic(err)
	}
}

func Copytree(srcpath, dstpath string) {
	sp := getpath(srcpath)
	dp := getpath(dstpath)
	if dp == sp {
		Panic("destination and source point to same path")
	}
	cptree(sp, dp)
}

func cptree(srcd, dstd string) {
	walk := func(path string, st os.FileInfo, err error) error {
		if err != nil {
			Panic(err)
		}
		var relp string
		relp, err = relpath(srcd, path)
		if err != nil {
			Panic(err)
		}
		dst := filepath.Join(dstd, relp)
		if st.IsDir() {
			if err := fs.MkdirAll(dst); err != nil {
				Panic(err)
			}
		} else if st.Mode().IsRegular() {
			cp(path, dst)
		} else {
			Log("WARN: copytree ignore non-regular file: %q", path)
		}
		return nil
	}
	filepath.Walk(srcd, walk)
}

func cp(src, dst string) {
	var err error
	var sfh *os.File
	var dfh *os.File
	if sfh, err = fs.Open(src); err != nil {
		Panic(err)
	}
	defer sfh.Close()
	if dfh, err = fs.Create(dst); err != nil {
		Panic(err)
	}
	defer dfh.Close()
	if err = fs.Copy(dfh, sfh); err != nil {
		Panic(err)
	}
	Log("cp: %q -> %q", src, dst)
}

func PathIsFile(path string) bool {
	path = getpath(path)
	st, err := fs.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			Panic(err)
		}
	}
	return st.Mode().IsRegular()
}

func Fread(filename string) *Content {
	filename = getpath(filename)
	blob, err := fs.ReadFile(filename)
	if err != nil {
		Panic(err)
	}
	return newContent(filename, blob)
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io"
	"os"
	"path/filepath"
)

var abspath func(string) (string, error) = filepath.Abs

func Rmtree(dpath string) {
	var err error
	d := filepath.FromSlash(dpath)
	d, err = abspath(d)
	if err != nil {
		Panic(err)
	}
	if st, err := os.Stat(d); err == nil {
		if st.IsDir() {
			Log("rmtree: %q", d)
		} else {
			Panicf("rmtree: %q is not a directory", d)
		}
	}
	if err := os.RemoveAll(d); err != nil {
		Panic(err)
	}
}

func Copytree(srcpath, dstpath string) {
	var err error
	sp := filepath.FromSlash(srcpath)
	sp, err = abspath(srcpath)
	if err != nil {
		Panic(err)
	}
	dp := filepath.FromSlash(dstpath)
	dp, err = abspath(dstpath)
	if err != nil {
		Panic(err)
	}
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
		relp, err = filepath.Rel(srcd, path)
		if err != nil {
			Panic(err)
		}
		dst := filepath.Join(dstd, relp)
		if st.IsDir() {
			if err := os.MkdirAll(dst, 0777); err != nil {
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
	if sfh, err = os.Open(src); err != nil {
		Panic(err)
	}
	defer sfh.Close()
	if dfh, err = os.Create(dst); err != nil {
		Panic(err)
	}
	defer dfh.Close()
	Log("cp: %q -> %q", src, dst)
	if _, err = io.Copy(dfh, sfh); err != nil {
		Panic(err)
	}
}

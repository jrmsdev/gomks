// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io"
	"os"
	"path/filepath"
)

func Rmtree(dpath string) {
	d := filepath.FromSlash(dpath)
	if err := os.RemoveAll(d); err != nil {
		Panic(err)
	}
}

func Copytree(srcpath, dstpath string) {
	var err error
	sp := filepath.FromSlash(srcpath)
	sp, err = filepath.Abs(srcpath)
	if err != nil {
		Panic(err)
	}
	dp := filepath.FromSlash(dstpath)
	dp, err = filepath.Abs(dstpath)
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
		if st.IsDir() {
			if err := os.MkdirAll(dstd, 0777); err != nil {
				Panic(err)
			}
			cptree(filepath.Join(srcd, path), filepath.Join(dstd, path))
		} else if st.Mode().IsRegular() {
			cp(filepath.Join(srcd, path), filepath.Join(dstd, path))
		} else {
			Log("copytree ignore non-regular file: %q", filepath.Join(srcd, path))
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
	if _, err = io.Copy(dfh, sfh); err != nil {
		Panic(err)
	}
}

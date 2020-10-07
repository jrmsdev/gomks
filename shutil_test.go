// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func lstree(t *testing.T, dpath string) []string {
	ls := make([]string, 0)
	walk := func(path string, st os.FileInfo, err error) error {
		if err != nil {
			t.Fatal(err)
		}
		if st.Mode().IsRegular() {
			relp, err := filepath.Rel(dpath, path)
			if err != nil {
				t.Fatal(err)
			}
			ls = append(ls, relp)
		}
		return nil
	}
	filepath.Walk(dpath, walk)
	return ls
}

func TestCopyRmtree(t *testing.T) {
	check := require.New(t)
	tmpdir, err := ioutil.TempDir("", "gomks.shutil_test")
	check.NoError(err)
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	t.Logf("tmpdir: %s", tmpdir)
	Copytree("./testdata/shutil/tree", filepath.ToSlash(filepath.Join(tmpdir, "shutil", "tree")))
	check.Equal(lstree(t, filepath.FromSlash("./testdata/shutil/tree")),
		lstree(t, filepath.Join(tmpdir, "shutil", "tree")))
	Rmtree(filepath.ToSlash(tmpdir))
}

func TestPathErrors(t *testing.T) {
	check := require.New(t)
	abspath = func(p string) (string, error) {
		if p == "mock/abspath/error" {
			return "", errors.New("mock error")
		}
		return p, nil
	}
	defer func() {
		abspath = filepath.Abs
	}()
	check.PanicsWithError("mock error", func() { Rmtree("mock/abspath/error") })
	check.PanicsWithError("mock error",
		func() { Copytree("mock/abspath/error", "fake/dest") })
	check.PanicsWithError("mock error",
		func() { Copytree("fake/source", "mock/abspath/error") })
	check.PanicsWithError("destination and source point to same path",
		func() { Copytree("fake/same", "fake/same") })
	check.PanicsWithError(`rmtree: "testdata/shutil/rmtree.txt" is not a directory`,
		func() { Rmtree("testdata/shutil/rmtree.txt") })
	relpath = func(p string, t string) (string, error) {
		return "", errors.New("mock error")
	}
	defer func() {
		relpath = filepath.Rel
	}()
	check.PanicsWithError("mock error", func() {
		cptree(filepath.FromSlash("./testdata/shutil/tree"),
			filepath.FromSlash("fake/dest"))
	})
}

func TestWalkError(t *testing.T) {
	check := require.New(t)
	check.Panics(func() { cptree("./testdata/notfound", "./testdata/_tmp") })
}

func TestRmtreeError(t *testing.T) {
	check := require.New(t)
	setMockFS("WithRemoveError")
	defer setNativeFS()
	check.PanicsWithError("mock remove error",
		func() { Rmtree("./testdata/shutil/rmtree") })
}

func TestCptreeError(t *testing.T) {
	check := require.New(t)
	setMockFS("WithMkdirError")
	defer setNativeFS()
	check.PanicsWithError("mock mkdir error", func() {
		cptree(filepath.FromSlash("./testdata/shutil/tree"),
			filepath.FromSlash("fake/dest")) })
	tmpdir := filepath.FromSlash("./testdata/_tmp")
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	cptree(os.DevNull, tmpdir)
}

func TestCpErrors(t *testing.T) {
	check := require.New(t)
	setMockFS("WithCopyError")
	defer setNativeFS()
	src := filepath.FromSlash("./testdata/shutil/tree/00.txt")
	dst := filepath.FromSlash("./testdata/_tmp")
	defer func() {
		os.RemoveAll(dst)
	}()
	check.PanicsWithError("mock copy error", func() { cp(src, dst) })
}

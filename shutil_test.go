// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func lstree(t *testing.T, dpath string) []string {
	t.Helper()
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
	defer setNativeFS()
	setMockFS()
	check.PanicsWithError(`rmtree: "testdata/shutil/rmtree.txt" is not a directory`, func() {
		Rmtree("testdata/shutil/rmtree.txt")
	})
	check.PanicsWithError("destination and source point to same path", func() {
		Copytree("testdata/shutil", "testdata/shutil")
	})
	setMockFS("WithPathError")
	check.PanicsWithError("mock relpath error", func() {
		cptree("testdata/shutil", "testdata/_tmp/shutil")
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
			filepath.FromSlash("fake/dest"))
	})
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
	setMockFS("WithCreateError")
	check.PanicsWithError("mock create error", func() { cp(src, dst) })
	setMockFS("WithOpenError")
	check.PanicsWithError("mock open error", func() { cp(src, dst) })
}

func TestPathIsFileError(t *testing.T) {
	check := require.New(t)
	setMockFS("WithStatError")
	defer setNativeFS()
	p := filepath.FromSlash("./testdata/shutil/tree/00.txt")
	check.PanicsWithError("mock stat error", func() { PathIsFile(p) })
}

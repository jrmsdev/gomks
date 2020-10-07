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

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

func TestCopyRmtree(t *testing.T) {
	check := require.New(t)
	tmpdir, err := ioutil.TempDir("", "gomks.shutil_test")
	check.NoError(err)
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	t.Logf("tmpdir: %s", tmpdir)
	Copytree("./testdata/shutil/tree", filepath.ToSlash(filepath.Join(tmpdir, "shutil", "tree")))
	Rmtree(tmpdir)
}

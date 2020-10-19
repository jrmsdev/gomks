// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTplBuild(t *testing.T) {
	check := require.New(t)
	vm := NewVM()
	err := vm.Eval(filepath.Join("testdata", "tpl", "build.mks"))
	check.NoError(err)
	diffCheck(t, filepath.Join("testdata", "tpl", "site"),
		filepath.Join("testdata", "_tmp", "tpl", "out"))
}

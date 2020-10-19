// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDemoBuild(t *testing.T) {
	check := require.New(t)
	vm := NewVM()
	demo := filepath.FromSlash("./demo")
	if err := os.Chdir(demo); err != nil {
		t.Fatal(err)
	}
	check.NoError(vm.Eval(filepath.FromSlash("./build.mks")))
	if err := os.Chdir(filepath.FromSlash("../")); err != nil {
		t.Fatal(err)
	}
	diffCheck(t, filepath.Join(demo, "_site"),
		filepath.Join("testdata", "demo", "_site"))
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	//~ "errors"
	//~ "io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDemoBuild(t *testing.T) {
	t.Skip()
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
	for _, fn := range lstree(t, filepath.Join(demo, "_site")) {
		src := filepath.Join(demo, "_site", fn)
		check.FileExists(src)
		chk := filepath.Join("testdata", "demo", "_site", fn)
		check.FileExists(chk)
		cmd := exec.Command("diff", "-Naur", chk, src)
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		check.NoError(err)
	}
}

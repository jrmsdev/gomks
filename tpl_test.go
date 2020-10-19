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

func TestTplRenderVfsErrors(t *testing.T) {
	check := require.New(t)
	tpl := TplParse("testdata/tpl/template/page.html")
	render := func() {
		TplRender("testdata/tpl/content/_index.html",
			"testdata/_tmp/tpl/site-index.html",
			tpl, ParamsNew())
	}
	defer setNativeFS()
	// write error
	setMockFS("WithWriteError")
	check.PanicsWithError("mock write error", render)
	// mkdir error
	setMockFS("WithMkdirError")
	check.PanicsWithError("mock mkdir error", render)
	// glob path error
	setMockFS("WithGlobError")
	check.PanicsWithError("mock glob error", render)
	// abs path error
	setMockFS("WithPathError")
	fs.(*mockFS).PathErrorCall = 2
	check.PanicsWithError("mock abspath error: 2", render)
}

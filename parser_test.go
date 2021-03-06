// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserDateSlug(t *testing.T) {
	check := require.New(t)
	fn := filepath.FromSlash("testdata/parser/index.html")
	c := readContent(fn)
	check.Equal("1970-01-01", c["date"])
	check.Equal("Thu, 01 Jan 1970 00:00:00 +0000", c["rfc_2822_date"])
	check.Equal("index", c["slug"])
	check.Equal("Index", c["title"])
	fn = filepath.FromSlash("testdata/parser/2020-10-08-date-slug.html")
	c = readContent(fn)
	check.Equal("2020-10-08", c["date"])
	check.Equal("Thu, 08 Oct 2020 00:00:00 +0000", c["rfc_2822_date"])
	check.Equal("date-slug", c["slug"])
	check.Equal("Date Slug", c["title"])
}

func TestParserReadHeaders(t *testing.T) {
	check := require.New(t)
	fn := filepath.FromSlash("testdata/parser/index.html")
	c := readContent(fn)
	check.Equal("Index", c["title"])
	check.IsType("", c["content"])
}

func TestParserReadContent(t *testing.T) {
	check := require.New(t)
	fn := filepath.FromSlash("testdata/parser/index.html")
	c := readContent(fn)
	check.Equal("Index", c["title"])
	keys := make([]string, 0)
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ck := []string{"slug", "title", "rfc_2822_date", "content", "date"}
	sort.Strings(ck)
	check.Equal(ck, keys)
}

func TestParserReadContentErrors(t *testing.T) {
	check := require.New(t)
	setMockFS("WithReadError")
	defer setNativeFS()
	// read error
	fn := filepath.FromSlash("testdata/parser/index.html")
	check.PanicsWithError("mock read error", func() { readContent(fn) })
	// time parse error
	setNativeFS()
	fn = filepath.FromSlash("testdata/parser/0000-00-00-index.html")
	check.Panics(func() { readContent(fn) })
}

func TestParserTemplatesLayout(t *testing.T) {
	check := require.New(t)
	params := ParamsNew()
	pageLayout := Fread("testdata/parser/layout/page.html")
	MakePages("testdata/parser/layout/_index.html",
		"testdata/_tmp/layout/index.html", pageLayout, params)
	blob, err := ioutil.ReadFile(filepath.FromSlash("testdata/_tmp/layout/index.html"))
	check.NoError(err)
	s := strings.Replace(string(blob), "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	check.Equal(`<!-- testing --><html><p>testing</p></html>`, s)
}

func TestParserMakePagesErrors(t *testing.T) {
	check := require.New(t)
	setMockFS("WithGlobError")
	defer setNativeFS()
	// glob error
	check.PanicsWithError("mock glob error", func() {
		MakePages("testdata/parser/index.html", "testdata/_tmp/_site", "", ParamsNew())
	})
	// abspath error
	setMockFS("WithPathError")
	fs.(*mockFS).PathErrorCall = 2
	check.PanicsWithError("mock abspath error: 2", func() {
		MakePages("testdata/parser/index.html", "testdata/_tmp/_site", "", ParamsNew())
	})
	// mkdir error
	setMockFS("WithMkdirError")
	check.PanicsWithError("mock mkdir error", func() {
		MakePages("testdata/parser/index.html", "testdata/_tmp/_site", "", ParamsNew())
	})
	// write error
	setMockFS("WithWriteError")
	check.PanicsWithError("mock write error", func() {
		MakePages("testdata/parser/index.html", "testdata/_tmp/_site", "", ParamsNew())
	})
}

func TestParserMakeListErrors(t *testing.T) {
	check := require.New(t)
	defer setNativeFS()
	// abspath error
	setMockFS("WithPathError")
	check.PanicsWithError("mock abspath error: 1", func() {
		MakeList(newPages(), "testdata/_tmp/index.html", "", "", ParamsNew())
	})
	// mkdir error
	setMockFS("WithMkdirError")
	check.PanicsWithError("mock mkdir error", func() {
		MakeList(newPages(), "testdata/_tmp/index.html", "", "", ParamsNew())
	})
	// write error
	setMockFS("WithWriteError")
	check.PanicsWithError("mock write error", func() {
		MakeList(newPages(), "testdata/_tmp/index.html", "", "", ParamsNew())
	})
}

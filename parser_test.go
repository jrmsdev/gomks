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
	check.Equal("Thu, 01 Jan 1970 00:00:00 +0000", c["rfc_date"])
	check.Equal("index", c["slug"])
	check.Equal("Index", c["title"])
	fn = filepath.FromSlash("testdata/parser/2020-10-08-date-slug.html")
	c = readContent(fn)
	check.Equal("2020-10-08", c["date"])
	check.Equal("Thu, 08 Oct 2020 00:00:00 +0000", c["rfc_date"])
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
	ck := []string{"slug", "title", "rfc_date", "content", "date"}
	sort.Strings(ck)
	check.Equal(ck, keys)
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

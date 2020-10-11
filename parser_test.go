// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderErrors(t *testing.T) {
	t.Log("TODO!!")
}

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
	check.IsType(template.HTML(""), c["content"])
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
	check.Equal(`<html><p>testing</p></html>`, s)
}

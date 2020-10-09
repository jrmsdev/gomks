// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderErrors(t *testing.T) {
	t.Log("TODO!!")
}

func TestReadContentDateSlug(t *testing.T) {
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

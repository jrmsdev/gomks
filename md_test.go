// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"io"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"

	"github.com/stretchr/testify/require"
)

func TestIsMarkdown(t *testing.T) {
	check := require.New(t)
	ok := []string{".md", ".mkd", ".mkdn", ".mdown", ".markdown"}
	for _, x := range ok {
		check.True(isMarkdown(x))
	}
	bad := []string{"", ".html", ".txt", ".xml"}
	for _, x := range bad {
		check.False(isMarkdown(x))
	}
}

func TestParseMarkdown(t *testing.T) {
	check := require.New(t)
	s := parseMarkdown([]byte("# testing"))
	check.Equal("<h1>testing</h1>\n", s)
}

func TestParseMarkdownError(t *testing.T) {
	check := require.New(t)
	mdConvert = func(src []byte, w io.Writer, opt ...parser.ParseOption) error {
		return errors.New("mock parser error")
	}
	defer func() {
		mdConvert = goldmark.Convert
	}()
	check.PanicsWithError("mock parser error", func() { parseMarkdown([]byte{}) })
}

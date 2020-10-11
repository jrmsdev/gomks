// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

var mdConvert func([]byte, io.Writer, ...parser.ParseOption) error = goldmark.Convert

func isMarkdown(ext string) bool {
	switch ext {
	case ".md":
		return true
	case ".mkd":
		return true
	case ".mkdn":
		return true
	case ".mdown":
		return true
	case ".markdown":
		return true
	}
	return false
}

func parseMarkdown(src []byte) string {
	buf := new(bytes.Buffer)
	if err := mdConvert(src, buf); err != nil {
		Panic(err)
	}
	defer buf.Reset()
	return buf.String()
}

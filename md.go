// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"

	"github.com/yuin/goldmark"
)

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
	if err := goldmark.Convert(src, buf); err != nil {
		Panic(err)
	}
	defer buf.Reset()
	return buf.String()
}

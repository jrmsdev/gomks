// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"html/template"
)

type Content struct {
	filename string
	data     []byte
}

func newContent(filename string, data []byte) *Content {
	return &Content{filename, data}
}

func (c *Content) Filename() string {
	return c.filename
}

func (c *Content) String() string {
	return string(c.data)
}

func (c *Content) HTML() template.HTML {
	return template.HTML(c.String())
}

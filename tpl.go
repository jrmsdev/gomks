// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"html/template"
	"path/filepath"
)

type Template struct {
	tpl *template.Template
}

func TplParse(pattern string) *Template {
	p := filepath.FromSlash(pattern)
	return &Template{tpl: template.Must(template.ParseGlob(p))}
}

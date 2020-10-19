// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"html/template"
)

type Template struct {
	tpl *template.Template
}

func TplLoad(pattern string) *Template {
	return &Template{tpl: template.Must(template.ParseGlob(pattern))}
}

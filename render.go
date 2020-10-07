// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
	"html/template"
)

func Render(tpl string, params paramMap) string {
	t, err := template.New("mks").Parse(tpl)
	if err != nil {
		Panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, params)
	if err != nil {
		Panic(err)
	}
	defer buf.Reset()
	return buf.String()
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
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

var tplExec func(*bytes.Buffer, *template.Template, interface{}) error

func tplExecImpl(buf *bytes.Buffer, tpl *template.Template, args interface{}) error {
	return tpl.Execute(buf, args)
}

func init() {
	tplExec = tplExecImpl
}

func TplRender(src, dst string, layout *Template, params paramMap) *Pages {
	tpl := template.Must(layout.tpl.Clone())
	src = filepath.FromSlash(src)
	flist, err := fs.Glob(src)
	if err != nil {
		Panic(err)
	}
	pages := newPages()
	dst = filepath.FromSlash(dst)
	buf := new(bytes.Buffer)
	for _, sp := range flist {
		c := readContent(sp)
		pages.Add(c)
		args := params.updateCopy(c)
		dp, err := fs.Abs(Render(dst, args))
		if err != nil {
			Panic(err)
		}
		args["content"] = template.HTML(args["content"].(string))
		Log("Render %q -> %q", sp, dp)
		if err := tplExec(buf, tpl, args); err != nil {
			Panic(err)
		}
		ddir := filepath.Dir(dp)
		if err := fs.MkdirAll(ddir); err != nil {
			Panic(err)
		}
		if err := fs.WriteFile(dp, buf.String()); err != nil {
			Panic(err)
		}
		buf.Reset()
	}
	pages.Sort()
	return pages
}

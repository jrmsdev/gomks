// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
	"html/template"
	"path/filepath"
	"regexp"
	"strings"
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

var reDateSlug *regexp.Regexp

func init() {
	reDateSlug = regexp.MustCompile(`^(?:(\d\d\d\d-\d\d-\d\d)-)?(.+)$`)
}

func readContent(fn string) paramMap {
	c := ParamsNew()
	_, err := fs.ReadFile(fn)
	if err != nil {
		Panic(err)
	}
	fn, err = filepath.Abs(fn)
	dateSlug := strings.Split(filepath.Base(fn), ".")[0]
	match := reDateSlug.FindStringSubmatch(dateSlug)
	c["date"] = match[1]
	if c["date"] == "" {
		c["date"] = "1970-01-01"
	}
	c["slug"] = match[2]
	return c
}

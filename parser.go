// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
	"encoding/json"
	"html/template"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reDateSlug *regexp.Regexp
	reHeader   *regexp.Regexp
)

func init() {
	reDateSlug = regexp.MustCompile(`^(?:(\d\d\d\d-\d\d-\d\d)-)?(.+)$`)
	reHeader = regexp.MustCompile(`^<!--\s*([^:]+):\s+(.*)\s*-->\r?\n`)
}

func Render(tpl *Content, params paramMap) string {
	t, err := template.New(tpl.Filename()).Parse(tpl.String())
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

type header struct {
	key string
	val interface{}
	end int
}

func readHeaders(text string) *header {
	match := reHeader.FindStringSubmatch(text)
	if len(match) == 0 {
		return &header{"", nil, -1}
	}
	h := &header{}
	h.key = match[1]
	if err := json.Unmarshal([]byte(match[2]), &h.val); err != nil {
		Panic(err)
	}
	h.end = len(match[0])
	return h
}

func readContent(fn string) paramMap {
	c := ParamsNew()
	// read file
	blob, err := fs.ReadFile(fn)
	if err != nil {
		Panic(err)
	}
	// get date and slug info
	dateSlug := strings.Split(filepath.Base(fn), ".")[0]
	match := reDateSlug.FindStringSubmatch(dateSlug)
	c["date"] = match[1]
	if c["date"] == "" {
		c["date"] = "1970-01-01"
	}
	c["slug"] = match[2]
	// read headers
	text := string(blob)
	for h := readHeaders(text); h.end > 0; h = readHeaders(text) {
		c[h.key] = h.val
		text = text[h.end:]
	}
	return c
}

func MakePages(src, dst string, layout *Content, params paramMap) {
}

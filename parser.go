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
	"time"
)

var (
	reDateSlug *regexp.Regexp
	reHeader   *regexp.Regexp
)

func init() {
	reDateSlug = regexp.MustCompile(`^(?:(\d\d\d\d-\d\d-\d\d)-)?(.+)$`)
	reHeader = regexp.MustCompile(`^<!--\s*([^\s]+):\s+(.*)\s*-->\r?\n`)
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

func readHeaders(fn string, blob []byte) *header {
	match := reHeader.FindSubmatch(blob)
	if len(match) == 0 {
		return &header{"", nil, -1}
	}
	h := &header{}
	h.key = string(match[1])
	if err := json.Unmarshal(match[2], &h.val); err != nil {
		Panicf("%s: %v", fn, err)
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
	var date time.Time
	if match[1] == "" {
		c["date"] = "1970-01-01"
		date = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		c["date"] = match[1]
		if date, err = time.Parse("2006-01-02", match[1]); err != nil {
			Panicf("%s: %v", fn, err)
		}
	}
	c["slug"] = match[2]
	// read headers
	for h := readHeaders(fn, blob); h.end > 0; h = readHeaders(fn, blob) {
		c[h.key] = h.val
		blob = blob[h.end:]
	}
	// TODO: convert markdown
	// update content
	c["rfc_date"] = date.Format(time.RFC1123Z)
	c["content"] = string(blob)
	return c
}

func MakePages(src, dst string, layout *Content, params paramMap) {
	src = filepath.FromSlash(src)
	flist, err := fs.Glob(src)
	if err != nil {
		Panic(err)
	}
	dst = filepath.FromSlash(dst)
	for _, sp := range flist {
		args := params.UpdateCopy(readContent(sp))
		dp := Render(&Content{"make_pages/dest_path", []byte(dst)}, args)
		dp, err = abspath(dp)
		if err != nil {
			Panic(err)
		}
		Log("Render %q -> %q", sp, dp)
		if err := fs.WriteFile(dp, Render(layout, args)); err != nil {
			Panic(err)
		}
	}
}

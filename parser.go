// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
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

func Render(tpl *Content, params paramMap) *Content {
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
	return newContent(tpl.Filename(), buf.Bytes())
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
	ext := filepath.Ext(filepath.Base(fn))
	if isMarkdown(ext) {
		c["content"] = parseMarkdown(blob)
	} else {
		c["content"] = string(blob)
	}
	return c
}

func MakePages(src, dst string, layout *Content, params paramMap) *Pages {
	src = filepath.FromSlash(src)
	flist, err := fs.Glob(src)
	if err != nil {
		Panic(err)
	}
	pages := newPages()
	dst = filepath.FromSlash(dst)
	for _, sp := range flist {
		c := readContent(sp)
		pages.Add(c)
		args := params.updateCopy(c)
		r := Render(&Content{"make_pages/dest_path", []byte(dst)}, args)
		dp, err := abspath(r.String())
		if err != nil {
			Panic(err)
		}
		Log("Render %q -> %q", sp, dp)
		ddir := filepath.Dir(dp)
		if err := fs.MkdirAll(ddir); err != nil {
			Panic(err)
		}
		r = Render(layout, args)
		if err := fs.WriteFile(dp, r.String()); err != nil {
			Panic(err)
		}
	}
	pages.Sort()
	return pages
}

func MakeList(pages *Pages, dst string, listLayout *Content, itemLayout *Content, params paramMap) {
	items := make([]string, 0)
	last := pages.len()
	for i := 0; i < last; i++ {
		p := params.updateCopy(pages.get(i))
		p["summary"] = "FIXME!!"
		r := Render(itemLayout, p)
		items = append(items, r.String())
	}
	r := Render(&Content{"make_list/dest_path", []byte(dst)}, params)
	dp, err := abspath(r.String())
	if err != nil {
		Panic(err)
	}
	params["content"] = strings.Join(items, "")
	Log("Render list %q", dp)
	ddir := filepath.Dir(dp)
	if err := fs.MkdirAll(ddir); err != nil {
		Panic(err)
	}
	r = Render(listLayout, params)
	if err := fs.WriteFile(dp, r.String()); err != nil {
		Panic(err)
	}
}

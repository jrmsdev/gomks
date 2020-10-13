// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	reDateSlug      *regexp.Regexp
	reHeader        *regexp.Regexp
	reTmpl          *regexp.Regexp
	reTruncateTags  *regexp.Regexp
	reTruncateSpace *regexp.Regexp
)

func init() {
	reDateSlug = regexp.MustCompile(`^(?:(\d\d\d\d-\d\d-\d\d)-)?(.+)$`)
	reHeader = regexp.MustCompile(`^<!--\s*([^\s]+):\s+(.*)\s+-->\r?\n`)
	reTmpl = regexp.MustCompile(`{{\s*([^}\s]+)\s*}}`)
	reTruncateTags = regexp.MustCompile(`(?s)<.*?>`)
	reTruncateSpace = regexp.MustCompile(`\r?\n`)
}

func Render(tpl string, params paramMap) string {
	return reTmpl.ReplaceAllStringFunc(tpl, func(s string) string {
		m := reTmpl.FindStringSubmatch(s)
		if v, found := params[m[1]]; found {
			return v.(string)
		}
		return s
	})
}

type header struct {
	key string
	val string
	end int
}

func readHeaders(fn string, blob []byte) *header {
	match := reHeader.FindSubmatch(blob)
	if len(match) == 0 {
		return &header{"", "", -1}
	}
	return &header{
		key: string(match[1]),
		val: string(match[2]),
		end: len(match[0]),
	}
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
	// update content
	c["rfc_2822_date"] = date.Format(time.RFC1123Z)
	ext := filepath.Ext(filepath.Base(fn))
	if isMarkdown(ext) {
		c["content"] = parseMarkdown(blob)
	} else {
		c["content"] = string(blob)
	}
	return c
}

func MakePages(src, dst string, layout string, params paramMap) *Pages {
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
		dp, err := fs.Abs(Render(dst, args))
		if err != nil {
			Panic(err)
		}
		Log("Render %q -> %q", sp, dp)
		ddir := filepath.Dir(dp)
		if err := fs.MkdirAll(ddir); err != nil {
			Panic(err)
		}
		if err := fs.WriteFile(dp, Render(layout, args)); err != nil {
			Panic(err)
		}
	}
	pages.Sort()
	return pages
}

func truncate(s string) string {
	s = reTruncateSpace.ReplaceAllString(s, " ")
	s = reTruncateTags.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	return strings.Join(strings.Split(s, " ")[:25], " ")
}

func MakeList(pages *Pages, dst string, listLayout string, itemLayout string, params paramMap) {
	items := make([]string, 0)
	last := pages.len()
	for i := 0; i < last; i++ {
		p := params.updateCopy(pages.get(i))
		p["summary"] = truncate(p["content"].(string))
		items = append(items, Render(itemLayout, p))
	}
	dp, err := fs.Abs(Render(dst, params))
	if err != nil {
		Panic(err)
	}
	params["content"] = strings.Join(items, "")
	Log("Render list %q", dp)
	ddir := filepath.Dir(dp)
	if err := fs.MkdirAll(ddir); err != nil {
		Panic(err)
	}
	if err := fs.WriteFile(dp, Render(listLayout, params)); err != nil {
		Panic(err)
	}
}

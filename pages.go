// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"sort"
	"strings"
)

var _ sort.Interface = &pSort{}

type pSort struct {
	p  *Pages
	by string
}

func newSortBy(p *Pages, by string) *pSort {
	return &pSort{p, by}
}

func (s *pSort) Len() int {
	return len(s.p.l)
}

func (s *pSort) Less(i, j int) bool {
	return strings.Compare(s.p.l[i][s.by].(string), s.p.l[j][s.by].(string)) == -1
}

func (s *pSort) Swap(i, j int) {
	jv := s.p.l[j]
	s.p.l[j] = s.p.l[i]
	s.p.l[i] = jv
}

type Pages struct {
	l []paramMap
}

func newPages() *Pages {
	return &Pages{l: make([]paramMap, 0)}
}

func (p *Pages) Add(c paramMap) {
	p.l = append(p.l, c)
}

func (p *Pages) Sort() {
	p.sortBy("date")
}

func (p *Pages) sortBy(key string) {
	sort.Sort(sort.Reverse(newSortBy(p, key)))
}

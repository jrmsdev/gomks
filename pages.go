// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

type Pages struct {
	l []paramMap
}

func newPages() *Pages {
	return &Pages{l: make([]paramMap, 0)}
}

func (p *Pages) Add(c paramMap) {
	p.l = append(p.l, c)
}

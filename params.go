// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"encoding/json"
	"path/filepath"
)

type paramMap map[string]interface{}

func ParamsNew() paramMap {
	return make(paramMap)
}

func (p paramMap) Load(filename string) {
	fn := filepath.FromSlash(filename)
	blob, err := fs.ReadFile(fn)
	if err != nil {
		Panic(err)
	}
	Log("Load params %q", fn)
	if err := json.Unmarshal(blob, &p); err != nil {
		Panic(err)
	}
}

func (p paramMap) Copy() paramMap {
	cp := ParamsNew()
	for k, v := range p {
		cp[k] = v
	}
	return cp
}

func (p paramMap) updateCopy(params paramMap) paramMap {
	cp := p.Copy()
	for k, v := range params {
		cp[k] = v
	}
	return cp
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var setenv func(string, string) error = os.Setenv

func SetenvDefault(key, val string) {
	if _, found := os.LookupEnv(key); !found {
		if err := setenv(key, val); err != nil {
			Panic(err)
		}
	}
}

func Getenv(key string) {
	os.Getenv(key)
}

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

func (p paramMap) UpdateCopy(params paramMap) paramMap {
	cp := p.Copy()
	for k, v := range params {
		cp[k] = v
	}
	return cp
}

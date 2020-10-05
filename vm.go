// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mks implements an scriptable static website generator.
package mks

import (
	"io/ioutil"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func Eval(e *env.Env, filename string) error {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if _, err := vm.Execute(e, nil, string(blob)); err != nil {
		return err
	}
	return nil
}

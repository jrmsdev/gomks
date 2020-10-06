// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io/ioutil"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

var _ VM = &vms{}

type VM interface {
	Execute(script string) error
	Eval(filename string) error
}

type vms struct {
	env  *env.Env
}

func NewVM() *vms {
	return &vms{env: newEnv()}
}

func (m *vms) Execute(script string) error {
	_, err := vm.Execute(m.env, nil, script)
	return err
}

func (m *vms) Eval(filename string) error {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := m.Execute(string(blob)); err != nil {
		return err
	}
	return nil
}

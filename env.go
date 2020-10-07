// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"github.com/mattn/anko/env"
)

var _ Env = &env.Env{}

type Env interface {
	Define(symbol string, value interface{}) error
}

func newEnv() *env.Env {
	e := env.NewEnv()
	define(e, "log", Log)
	define(e, "version", Version)
	define(e, "panic", Panic)
	define(e, "rmtree", Rmtree)
	define(e, "copytree", Copytree)
	define(e, "args_new", ArgsNew)
	define(e, "params_new", ParamsNew)
	return e
}

func define(e Env, symbol string, value interface{}) {
	if err := e.Define(symbol, value); err != nil {
		panic(err)
	}
}

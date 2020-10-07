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
	define(e, "params_new", ParamsNew)
	define(e, "setenv_default", SetenvDefault)
	define(e, "getenv", Getenv)
	define(e, "path_isfile", PathIsFile)
	return e
}

func define(e Env, symbol string, value interface{}) {
	if err := e.Define(symbol, value); err != nil {
		panic(err)
	}
}

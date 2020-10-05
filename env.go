// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

import (
	"github.com/mattn/anko/env"
)

type Env struct {
	*env.Env
}

func NewEnv() *Env {
	e := env.NewEnv()
	define(e, "log", Log)
	define(e, "version", Version)
	return &Env{Env: e}
}

func define(e *env.Env, symbol string, value interface{}) {
	if err := e.Define(symbol, value); err != nil {
		panic(err)
	}
}

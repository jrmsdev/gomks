// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mks, an static website generator.
package mks

import (
	"log"

	"github.com/mattn/anko/env"
	//~ "github.com/mattn/anko/vm"
)

const version string = "0.0"

func Version() string {
	return version
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

type Env struct {
	*env.Env
}

func NewEnv() *Env {
	e := env.NewEnv()
	e.Define("log", Log)
	return &Env{Env: e}
}

func Eval(e *Env, filename string) error {
	return nil
}

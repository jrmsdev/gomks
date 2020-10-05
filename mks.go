// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mks, an static website generator.
package mks

import (
	"io/ioutil"
	"log"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

var version string = "master"

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
	define(e, "log", Log)
	define(e, "version", Version)
	return &Env{Env: e}
}

func define(e *env.Env, symbol string, value interface{}) {
	if err := e.Define(symbol, value); err != nil {
		panic(err)
	}
}

func Eval(e *Env, filename string) error {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if _, err := vm.Execute(e.Env, nil, string(blob)); err != nil {
		return err
	}
	return nil
}

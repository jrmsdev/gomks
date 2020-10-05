// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package main implements the mks command line tool.
package main

import (
	"os"

	"github.com/jrmsdev/gomks"
)

func main() {
	if len(os.Args) != 2 {
		mks.Log("ERROR: %s", "invalid args")
		os.Exit(1)
	}
	if os.Args[1] == "version" {
		mks.Log("mks version %s", mks.Version())
		os.Exit(0)
	}
	e := mks.NewEnv()
	for _, fn := range os.Args[1:] {
		if err := mks.Eval(e, fn); err != nil {
			mks.Log("ERROR: %s: %v", fn, err)
			os.Exit(2)
		}
	}
	os.Exit(0)
}

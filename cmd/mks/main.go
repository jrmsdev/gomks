// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package main implements the mks command line tool.
package main

import (
	"flag"
	"os"

	mks "github.com/jrmsdev/gomks"
)

var exit func(int) = os.Exit

func main() {
	flag.Parse()
	exit(run(flag.Args()))
}

func run(args []string) int {
	if len(args) == 0 {
		mks.Log("ERROR: %s", "no args")
		return 1
	}
	if args[0] == "version" {
		mks.Log("mks version %s", mks.Version())
		return 0
	}
	e := mks.NewEnv()
	for _, fn := range args {
		if err := mks.Eval(e, fn); err != nil {
			mks.Log("ERROR: %s: %v", fn, err)
			return 2
		}
	}
	return 0
}

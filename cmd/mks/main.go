// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package main, the mks command line tool.
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
	arg := os.Args[1]
	switch arg {
	case "version":
		mks.Log("mks version %s", mks.Version())
	default:
		mks.Log("ERROR: invalid action %q", arg)
		os.Exit(2)
	}
	os.Exit(0)
}

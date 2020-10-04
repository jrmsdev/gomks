// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package main, the mks command line tool.
package main

import (
	"os"

	"github.com/jrmsdev/gomks"
)

func main() {
	mks.Log("Args: %#v", os.Args)
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mks, an static website generator.
package mks

import (
	"log"
)

func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package gomks implements an scriptable static website generator.
package gomks

import (
	"log"
)

var version string = "master"

func Version() string {
	return version
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

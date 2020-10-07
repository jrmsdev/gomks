// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package gomks implements an scriptable static website generator.
package gomks

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lmsgprefix)
	if os.Getenv("MKSLOG") == "debug" {
		log.SetFlags(log.Llongfile)
	}
}

var version string = "master"

func Version() string {
	return version
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Panic(v interface{}) {
	log.Output(2, fmt.Sprint(v))
	panic(v)
}

func Panicf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Output(2, msg)
	panic(msg)
}

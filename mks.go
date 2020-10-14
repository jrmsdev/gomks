// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package gomks implements an scriptable static website generator.
package gomks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(0)
	if os.Getenv("MKSLOG") == "debug" {
		log.SetFlags(log.Llongfile)
	}
}

var version string = "0.3"
var build string = ""

func Version() string {
	if build == "" {
		return version
	}
	return build[1:]
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Panic(v interface{}) {
	msg := fmt.Sprint(v)
	log.Output(2, msg)
	panic(errors.New(msg))
}

func Panicf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Output(2, msg)
	panic(errors.New(msg))
}

var setenv func(string, string) error = os.Setenv

func SetenvDefault(key, val string) {
	if _, found := os.LookupEnv(key); !found {
		if err := setenv(key, val); err != nil {
			Panic(err)
		}
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}

func Datetime() time.Time {
	return time.Now().Local()
}

func DatetimeUTC() time.Time {
	return time.Now().UTC()
}

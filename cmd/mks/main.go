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

var (
	showVersion bool
	serveSite   string
	httpListen  string
)

func main() {
	flag.BoolVar(&showVersion, "version", false, "show version and exit")
	flag.StringVar(&serveSite, "serve", "", "run http server on `site_dirpath`")
	flag.StringVar(&httpListen, "http", "127.0.0.1:8080",
		"bind http server to `address:port`")
	flag.Parse()
	exit(run(flag.Args()))
}

func run(args []string) int {
	if showVersion {
		mks.Log("mks version %s", mks.Version())
		return 0
	}
	if serveSite != "" {
		return runServer(httpListen, serveSite)
	}
	if len(args) == 0 {
		mks.Log("ERROR: %s", "no args")
		return 1
	}
	vm := mks.NewVM()
	for _, fn := range args {
		if err := vm.Eval(fn); err != nil {
			mks.Log("ERROR: %s: %v", fn, err)
			return 2
		}
	}
	return 0
}

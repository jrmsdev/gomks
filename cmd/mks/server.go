// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"net/http"

	mks "github.com/jrmsdev/gomks"
)

func runServer(addr, site string) int {
	mks.Log("Listen http://%s", addr)
	err := http.ListenAndServe(addr, http.FileServer(http.Dir(site)))
	if err != nil {
		mks.Log("ERROR: %v", err)
		return 9
	}
	return 0
}

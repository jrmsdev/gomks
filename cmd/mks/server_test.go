// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	check := require.New(t)
	lns = func(addr string, h http.Handler) error {
		return nil
	}
	serveSite = filepath.FromSlash("testdata/_site")
	defer func() {
		lns = http.ListenAndServe
		serveSite = ""
	}()
	rc := run([]string{})
	check.Equal(0, rc)
}

func TestServerError(t *testing.T) {
	check := require.New(t)
	lns = func(addr string, h http.Handler) error {
		return errors.New("mock server error")
	}
	defer func() {
		lns = http.ListenAndServe
	}()
	rc := runServer("testing", "server")
	check.Equal(9, rc)
}

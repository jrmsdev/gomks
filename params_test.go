// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamsNew(t *testing.T) {
	check := require.New(t)
	p := ParamsNew()
	check.IsType(paramMap{}, p)
}

func TestParams(t *testing.T) {
	check := require.New(t)
	p := ParamsNew()
	p["test"] = "ing"
	check.Equal("ing", p["test"])
}

func TestParamsLoad(t *testing.T) {
	check := require.New(t)
	p := ParamsNew()
	p["test"] = "ing"
	check.Equal("ing", p["test"])
	p.Load("testdata/params/update.json")
	check.Equal("testing", p["test"])
}

func TestParamsLoadError(t *testing.T) {
	check := require.New(t)
	p := ParamsNew()
	check.PanicsWithError("invalid character '}' looking for beginning of value",
		func() { p.Load("testdata/params/update-error.json") })
	setMockFS("WithReadError")
	defer setNativeFS()
	check.PanicsWithError("mock read error",
		func() { p.Load("testdata/params/update-error.json") })
}

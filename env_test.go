// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"testing"

	"github.com/mattn/anko/env"
	"github.com/stretchr/testify/suite"
)

func TestEnv(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

type EnvTestSuite struct {
	suite.Suite
}

func (s *EnvTestSuite) SetupTest() {
}

func (s *EnvTestSuite) TearDownTest() {
}

type symt struct {
	script string
	err    string
}

var symTests = map[string]map[string]symt{
	"log": {
		"test": {`log("test%s", "ing")`, ""},
	},
	"version": {
		"test": {"version()", ""},
	},
	"panic": {
		"test": {`panic("testing")`, "testing"},
	},
}

func getSym(e *env.Env, n string) error {
	_, err := e.Get(n)
	return err
}

func (s *EnvTestSuite) TestSymbols() {
	check := s.Require()
	for n := range symTests {
		e := newEnv()
		check.NoError(getSym(e, n), "get symbol %q", n)
		for nn := range symTests[n] {
			vm := NewVM()
			t := symTests[n][nn]
			if t.err != "" {
				check.EqualError(vm.Execute(t.script), t.err, "symbol %q test %q", n, nn)
			} else {
				check.NoError(vm.Execute(t.script), t.err, "symbol %q test %q", n, nn)
			}
		}
	}
}

var _ Env = &mockEnv{}

type mockEnv struct {
	err error
}

func (e *mockEnv) Define(symbol string, value interface{}) error {
	return e.err
}

func (s *EnvTestSuite) TestDefinePanic() {
	check := s.Require()
	e := new(mockEnv)
	e.err = errors.New("testing")
	x := func() {
		define(e, "nosym", nil)
	}
	check.PanicsWithError("testing", x)
}

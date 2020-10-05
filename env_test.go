// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

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

func getSym(e *env.Env, n string) error {
	_, err := e.Get(n)
	return err
}

func (s *EnvTestSuite) TestSymbols() {
	check := s.Require()
	e := NewEnv()
	check.NoError(getSym(e, "log"))
	check.NoError(getSym(e, "version"))
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

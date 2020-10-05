// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

import (
	"testing"

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

func getSym(e *Env, n string) error {
	_, err := e.Env.Get(n)
	return err
}

func (s *EnvTestSuite) TestSymbols() {
	check := s.Require()
	e := NewEnv()
	check.NoError(getSym(e, "log"))
	check.NoError(getSym(e, "version"))
}

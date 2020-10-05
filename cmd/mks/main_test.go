// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

type MainTestSuite struct {
	suite.Suite
	ExitCode int
}

func (s *MainTestSuite) SetupTest() {
	s.ExitCode = -1
	exit = s.mockExit
}

func (s *MainTestSuite) TearDownTest() {
	exit = os.Exit
}

func (s *MainTestSuite) mockExit(code int) {
	s.ExitCode = code
}

func (s *MainTestSuite) TestNoArgs() {
	check := s.Require()
	check.Equal(-1, s.ExitCode)
	main()
	check.Equal(1, s.ExitCode)
}

func (s *MainTestSuite) TestVersion() {
	check := s.Require()
	rc := run([]string{"version"})
	check.Equal(0, rc)
}

func (s *MainTestSuite) TestEval() {
	check := s.Require()
	rc := run([]string{"testdata/eval.mks"})
	check.Equal(0, rc)
}

func (s *MainTestSuite) TestEvalError() {
	check := s.Require()
	rc := run([]string{"eval.error"})
	check.Equal(2, rc)
}

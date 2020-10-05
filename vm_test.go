// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

import (
	//~ "errors"
	"testing"

	//~ "github.com/mattn/anko/env"
	"github.com/stretchr/testify/suite"
)

func TestVM(t *testing.T) {
	suite.Run(t, new(VMTestSuite))
}

type VMTestSuite struct {
	suite.Suite
}

func (s *VMTestSuite) SetupTest() {
}

func (s *VMTestSuite) TearDownTest() {
}

func (s *VMTestSuite) TestNewVM() {
	check := s.Require()
	vm := NewVM()
	check.NotNil(vm.opts)
	check.False(vm.opts.Debug)
	check.NotNil(vm.env)
}

func (s *VMTestSuite) TestExec() {
	check := s.Require()
	vm := NewVM()
	check.NoError(vm.Execute(`log("test%s", "ing")`))
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

import (
	//~ "errors"
	"path/filepath"
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

func (s *VMTestSuite) TestEval() {
	check := s.Require()
	vm := NewVM()
	check.NoError(vm.Eval(filepath.FromSlash("./testdata/eval.mks")))
}

func (s *VMTestSuite) TestEvalError() {
	check := s.Require()
	vm := NewVM()
	check.Error(vm.Eval("eval.error"))
}

func (s *VMTestSuite) TestExecError() {
	check := s.Require()
	vm := NewVM()
	check.EqualError(vm.Eval(filepath.FromSlash("./testdata/eval-error.mks")), "undefined symbol 'invalid'")
}

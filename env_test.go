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
	setMockFS()
}

func (s *EnvTestSuite) TearDownTest() {
	setNativeFS()
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
	"copytree": {
		"test":  {`copytree("testdata/shutil/tree", "testdata/_tmp/shutil/tree")`, ""},
		"clean": {`rmtree("testdata/_tmp")`, ""},
	},
	"rmtree": {
		"copy": {`copytree("testdata/shutil/tree", "testdata/_tmp/shutil/tree")`, ""},
		"test": {`rmtree("testdata/_tmp")`, ""},
	},
	"params_new": {
		"test": {"args = params_new()", ""},
	},
	"setenv_default": {
		"test": {`setenv_default("TESTING", "gomks")`, ""},
	},
	"getenv": {
		"test": {`getenv("TESTING")`, ""},
	},
	"path_isfile": {
		"test":     {`path_isfile("testdata/shutil")`, ""},
		"notfound": {`path_isfile("testdata/shutil/found.not")`, ""},
	},
	"fread": {
		"test": {`fread("testdata/shutil/tree/00.txt")`, ""},
		"notfound": {`fread("testdata/shutil/found.not")`,
			"open testdata/shutil/found.not: no such file or directory"},
	},
	"render": {
		"test": {`render(fread("testdata/render/test.html"), params_new())`, ""},
	},
	"make_pages": {
		"test": {`make_pages("", "",
			fread("testdata/render/test.html"), params_new())`, ""},
	},
	"make_list": {
		"test": {`make_list(make_pages("", "",
			fread("testdata/render/test.html"), params_new()),
			"testdata/render/test.html", "", "", params_new())`, ""},
	},
	"datetime": {
		"test": {"datetime().Year()", ""},
	},
	"datetime_utc": {
		"test": {"datetime_utc()", ""},
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
				check.EqualError(vm.Execute(t.script), t.err, "symbol %q %q", n, nn)
			} else {
				check.NoError(vm.Execute(t.script), "symbol %q %q", n, nn)
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

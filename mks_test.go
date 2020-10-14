// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	check := assert.New(t)
	check.Equal(version, Version())
	check.Equal("", build)
	build = "vtesting"
	defer func() {
		build = ""
	}()
	check.Equal("testing", Version())
}

func TestLog(t *testing.T) {
	Log("test%s", "ing")
}

func TestPanic(t *testing.T) {
	check := assert.New(t)
	f := func() {
		Panic("testing")
	}
	check.PanicsWithError("testing", f)
}

func TestPanicf(t *testing.T) {
	check := assert.New(t)
	f := func() {
		Panicf("test%s", "ing")
	}
	check.PanicsWithError("testing", f)
}

func TestSetGetEnv(t *testing.T) {
	check := require.New(t)
	n := fmt.Sprintf("GOMKS_ENV_TEST_%d_%d", rand.Int(), rand.Int())
	check.Equal("", Getenv(n))
	SetenvDefault(n, "testing")
	check.Equal("testing", Getenv(n))
	SetenvDefault(n, "test2")
	check.Equal("testing", Getenv(n))
}

func TestSetenvError(t *testing.T) {
	check := require.New(t)
	setenv = func(k, v string) error {
		return errors.New("mock error")
	}
	defer func() {
		setenv = os.Setenv
	}()
	check.PanicsWithError("mock error", func() { SetenvDefault("TEST", "ing") })
}

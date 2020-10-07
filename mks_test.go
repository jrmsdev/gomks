// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	check := assert.New(t)
	check.Equal(version, Version())
}

func TestLog(t *testing.T) {
	Log("test%s", "ing")
}

func TestPanic(t *testing.T) {
	check := assert.New(t)
	f := func() {
		Panic("testing")
	}
	check.PanicsWithValue("testing", f)
}

func TestPanicf(t *testing.T) {
	check := assert.New(t)
	f := func() {
		Panicf("test%s", "ing")
	}
	check.PanicsWithValue("testing", f)
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	check := assert.New(t)
	check.Equal(fmt.Sprintf("%s-%s", version, build), Version())
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

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mks

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

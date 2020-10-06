// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRmtree(t *testing.T) {
	check := assert.New(t)
	check.Equal(version, Version())
}

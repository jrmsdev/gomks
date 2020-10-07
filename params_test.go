// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgsNew(t *testing.T) {
	check := require.New(t)
	check.IsType(map[string]string{}, ArgsNew())
}

func TestParamsNew(t *testing.T) {
	check := require.New(t)
	check.IsType(map[string]string{}, ParamsNew())
}

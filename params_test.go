// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamsNew(t *testing.T) {
	check := require.New(t)
	check.IsType(map[string]string{}, ParamsNew())
}

func TestSetenvError(t *testing.T) {
	check := require.New(t)
	setenv = func(k, v string) error {
		return errors.New("mock error")
	}
	defer func() {
		setenv = os.Setenv
	}()
	check.PanicsWithError("mock error", func() { SetenvDefault("TEST", "ing") } )
}

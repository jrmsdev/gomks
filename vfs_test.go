// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
)

type mockFS struct {
	fs              fsi
	WithRemoveError bool
}

func setMockFS(args ...string) {
	m := &mockFS{fs: &nativeFS{}}
	for _, a := range args {
		switch a {
		case "WithRemoveError":
			m.WithRemoveError = true
		}
	}
	fs = nil
	fs = m
}

func (m *mockFS) RemoveAll(p string) error {
	if m.WithRemoveError {
		return errors.New("mock remove error")
	}
	return m.fs.RemoveAll(p)
}

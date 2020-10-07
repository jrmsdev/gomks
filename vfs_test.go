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

func setMockFS() {
	fs = nil
	fs = &mockFS{fs: &nativeFS{}}
}

func (m *mockFS) RemoveAll(p string) error {
	if m.WithRemoveError {
		return errors.New("mock remove error")
	}
	return m.fs.RemoveAll(p)
}

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"io"
	"os"
)

type mockFS struct {
	fs              fsi
	WithRemoveError bool
	WithMkdirError  bool
	WithCopyError   bool
	WithCreateError bool
}

func setMockFS(args ...string) {
	m := &mockFS{fs: &nativeFS{}}
	for _, a := range args {
		switch a {
		case "WithRemoveError":
			m.WithRemoveError = true
		case "WithMkdirError":
			m.WithMkdirError = true
		case "WithCopyError":
			m.WithCopyError = true
		case "WithCreateError":
			m.WithCreateError = true
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

func (m *mockFS) MkdirAll(p string) error {
	if m.WithMkdirError {
		return errors.New("mock mkdir error")
	}
	return m.fs.MkdirAll(p)
}

func (m *mockFS) Copy(dst io.Writer, src io.Reader) error {
	if m.WithCopyError {
		return errors.New("mock copy error")
	}
	return m.fs.Copy(dst, src)
}

func (m *mockFS) Create(p string) (*os.File, error) {
	if m.WithCreateError {
		return nil, errors.New("mock create error")
	}
	return m.fs.Create(p)
}

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
	WithOpenError   bool
	WithStatError   bool
	WithReadError   bool
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
		case "WithOpenError":
			m.WithOpenError = true
		case "WithStatError":
			m.WithStatError = true
		case "WithReadError":
			m.WithReadError = true
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

func (m *mockFS) Open(p string) (*os.File, error) {
	if m.WithOpenError {
		return nil, errors.New("mock open error")
	}
	return m.fs.Open(p)
}

func (m *mockFS) Stat(p string) (os.FileInfo, error) {
	if m.WithStatError {
		return nil, errors.New("mock stat error")
	}
	return m.fs.Stat(p)
}

func (m *mockFS) ReadAll(r io.Reader) ([]byte, error) {
	if m.WithReadError {
		return nil, errors.New("mock read error")
	}
	return m.fs.ReadAll(r)
}

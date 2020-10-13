// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"io"
	"os"
	"path/filepath"
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
	WithGlobError   bool
	WithWriteError  bool
	WithPathError   bool
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
		case "WithGlobError":
			m.WithGlobError = true
		case "WithWriteError":
			m.WithWriteError = true
		case "WithPathError":
			m.WithPathError = true
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

func (m *mockFS) ReadFile(p string) ([]byte, error) {
	if m.WithReadError {
		return nil, errors.New("mock read error")
	}
	return m.fs.ReadFile(p)
}

func (m *mockFS) Glob(p string) ([]string, error) {
	if m.WithGlobError {
		return nil, errors.New("mock glob error")
	}
	return m.fs.Glob(p)
}

func (m *mockFS) WriteFile(p string, b string) error {
	if m.WithWriteError {
		return errors.New("mock write error")
	}
	return m.fs.WriteFile(p, b)
}

func (m *mockFS) Abs(p string) (string, error) {
	if m.WithPathError {
		return "", errors.New("mock abspath error")
	}
	return m.fs.Abs(p)
}

func (m *mockFS) Rel(b, p string) (string, error) {
	if m.WithPathError {
		return "", errors.New("mock relpath error")
	}
	return m.fs.Rel(b, p)
}

func (m *mockFS) GetPath(s string) string {
	return filepath.FromSlash(s)
}

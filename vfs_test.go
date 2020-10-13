// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVFSErrors(t *testing.T) {
	check := require.New(t)
	setMockFS("WithPathError")
	defer setNativeFS()
	n := newNativeFS()
	check.PanicsWithError("mock abspath error: 1", func() {
		n.GetPath("testdata/render/test.html")
	})
	// abspath error
	setMockFS("WithPathError")
	_, err := n.Glob("testdata/shutil/tree/*.txt")
	check.EqualError(err, "mock abspath error: 1")
	// glob error
	setMockFS()
	n.glob = func(p string) ([]string, error) {
		return nil, errors.New("mock glob error")
	}
	_, err = n.Glob("testing")
	check.EqualError(err, "mock glob error")
}

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
	PathErrorCall   int
	pcall           int
}

func setMockFS(args ...string) {
	m := &mockFS{fs: newNativeFS()}
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
	m.PathErrorCall = 1
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
	m.pcall += 1
	if m.WithPathError {
		if m.pcall == m.PathErrorCall {
			return "", errors.New(fmt.Sprintf("mock abspath error: %d", m.pcall))
		}
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

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var fs fsi

type fsi interface {
	RemoveAll(string) error
	MkdirAll(string) error
	Copy(io.Writer, io.Reader) error
	Create(string) (*os.File, error)
	Open(string) (*os.File, error)
	Stat(string) (os.FileInfo, error)
	ReadFile(string) ([]byte, error)
	Glob(string) ([]string, error)
	WriteFile(string, string) error
	Abs(string) (string, error)
	Rel(string, string) (string, error)
	GetPath(s string) string
}

type nativeFS struct {
}

func setNativeFS() {
	fs = nil
	fs = &nativeFS{}
}

func init() {
	setNativeFS()
}

func (n *nativeFS) RemoveAll(p string) error {
	return os.RemoveAll(p)
}

func (n *nativeFS) MkdirAll(p string) error {
	return os.MkdirAll(p, 0777)
}

func (n *nativeFS) Copy(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}

func (n *nativeFS) Create(p string) (*os.File, error) {
	return os.Create(p)
}

func (n *nativeFS) Open(p string) (*os.File, error) {
	return os.Open(p)
}

func (n *nativeFS) Stat(p string) (os.FileInfo, error) {
	return os.Stat(p)
}

func (n *nativeFS) ReadFile(p string) ([]byte, error) {
	return ioutil.ReadFile(p)
}

func (n *nativeFS) Glob(p string) ([]string, error) {
	l, err := filepath.Glob(p)
	if err != nil {
		return nil, err
	}
	flist := make([]string, 0)
	for _, fn := range l {
		n, err := filepath.Abs(fn)
		if err != nil {
			return nil, err
		}
		flist = append(flist, n)
	}
	return flist, nil
}

func (n *nativeFS) WriteFile(p string, b string) error {
	return ioutil.WriteFile(p, []byte(b), 0666)
}

func (n *nativeFS) Abs(p string) (string, error) {
	return filepath.Abs(p)
}

func (n *nativeFS) Rel(b, p string) (string, error) {
	return filepath.Rel(b, p)
}

func (n *nativeFS) GetPath(s string) string {
	var err error
	p := filepath.FromSlash(s)
	p, err = n.Abs(p)
	if err != nil {
		Panic(err)
	}
	return p
}

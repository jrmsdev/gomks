// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"io"
	"io/ioutil"
	"os"
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

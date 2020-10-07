// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
)

var fs fsi

type fsi interface {
	RemoveAll(string) error
	MkdirAll(string) error
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

// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
)

var fs fsi

type fsi interface {
	RemoveAll(string) error
}

type nativeFS struct {
}

func setNativeFS() {
	fs = nil
	fs = &nativeFS{}
}

func (n *nativeFS) RemoveAll(p string) error {
	return os.RemoveAll(p)
}

func init() {
	setNativeFS()
}

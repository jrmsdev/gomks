// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
)

var fs fsi = &nativeFS{}

type fsi interface {
	RemoveAll(string) error
}

type nativeFS struct {
}

func (n *nativeFS) RemoveAll(p string) error {
	return os.RemoveAll(p)
}

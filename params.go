// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package gomks

import (
	"os"
)

var setenv func(string, string) error = os.Setenv

func SetenvDefault(key, val string) {
	if _, found := os.LookupEnv(key); !found {
		if err := setenv(key, val); err != nil {
			Panic(err)
		}
	}
}

func Getenv(key string) {
	os.Getenv(key)
}

func ParamsNew() map[string]string {
	return make(map[string]string)
}

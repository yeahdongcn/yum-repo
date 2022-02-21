//go:build !linux
// +build !linux

package repo

import "errors"

func doWatch(path string, done chan struct{}) error {
	return errors.New("watcher unsupported in this build")
}

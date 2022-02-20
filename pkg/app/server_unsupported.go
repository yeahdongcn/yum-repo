//go:build !linux
// +build !linux

package app

import "errors"

func watchForLockfileContention(path string, done chan struct{}) error {
	return errors.New("kubelet unsupported in this build")
}

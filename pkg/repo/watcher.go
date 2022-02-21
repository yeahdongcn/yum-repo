package repo

import (
	"k8s.io/klog/v2"
)

func Watch(repo string) error {
	done := make(chan struct{})
	if err := doWatch(repo, done); err != nil {
		klog.ErrorS(err, "Unable to watch")
	}
	<-done
	return nil
}

package app

import (
	"k8s.io/klog/v2"
)

func X() {
	done := make(chan struct{})
	if err := watchForLockfileContention("", done); err != nil {
		klog.ErrorS(err, "Unable to watch")
	}
}

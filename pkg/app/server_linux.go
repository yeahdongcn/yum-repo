package app

import (
	"k8s.io/klog/v2"
	"k8s.io/utils/inotify"
)

func watchForLockfileContention(path string, done chan struct{}) error {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		klog.ErrorS(err, "Unable to create watcher for lockfile")
		return err
	}
	if err = watcher.AddWatch(path, inotify.InOpen|inotify.InDeleteSelf); err != nil {
		klog.ErrorS(err, "Unable to watch lockfile")
		watcher.Close()
		return err
	}
	go func() {
		select {
		case ev := <-watcher.Event:
			klog.InfoS("inotify event", "event", ev)
		case err = <-watcher.Error:
			klog.ErrorS(err, "inotify watcher error")
		}
		close(done)
		watcher.Close()
	}()
	return nil
}

package repo

import (
	"errors"

	"k8s.io/klog/v2"
	"k8s.io/utils/inotify"
)

func doneWatch(watcher *inotify.Watcher, done chan struct{}) {
	if watcher != nil {
		watcher.Close()
	}
	close(done)
}

func doWatch(path string, done chan struct{}) error {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		klog.ErrorS(err, "Unable to create watcher for repo")
		doneWatch(watcher, done)
		return err
	}
	flags := inotify.InCreate |
		inotify.InDelete | inotify.InDeleteSelf |
		inotify.InModify |
		inotify.InMove | inotify.InMoveSelf |
		inotify.InIgnored
	if err = watcher.AddWatch(path, flags); err != nil {
		klog.ErrorS(err, "Unable to watch repo", "path", path)
		doneWatch(watcher, done)
		return err
	}
	klog.Info("Watch started")
	go func() {
		var retErr error
		for {
			select {
			case ev := <-watcher.Event:
				klog.InfoS("inotify event", "event", ev)
				if ev.Mask == inotify.InIgnored {
					retErr = errors.New("watch was removed explicitly or automatically")
				} else {
					// TODO:
				}
			case err = <-watcher.Error:
				klog.ErrorS(err, "inotify watcher error")
				retErr = err
			}
			if retErr != nil {
				break
			}
		}
		doneWatch(watcher, done)
	}()
	return nil
}

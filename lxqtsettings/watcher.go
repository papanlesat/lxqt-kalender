package lxqtsettings

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func (a *Applier) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Failed to initialize file watcher:", err)
		return
	}
	defer watcher.Close()

	// Watch LXQt config directory
	err = watcher.Add(a.configDir)
	if err != nil {
		log.Println("Failed to watch LXQt config directory:", err)
		return
	}

	log.Println("Watching", a.configDir, "for changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
				ext := filepath.Ext(event.Name)
				if ext == ".conf" {
					log.Printf("[LXQtSettings] Detected change in %s, reapplying settings...\n", event.Name)
					a.Apply()
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

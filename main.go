package main

import (
	"github.com/radovskyb/watcher"
	"log"
	"os"
	"time"
)

func main() {
	dirname := getGtdDir()

	log.Println("Watching directory: ", dirname)

	fileWatcher := watcher.New()
	fileWatcher.Add(dirname)

	go func() {
		for {
			select {
			case _ = <-fileWatcher.Event:
				OrganizeGtdFolder(dirname)
			case err := <-fileWatcher.Error:
				log.Fatalln(err)
			case <-fileWatcher.Closed:
				return
			}
		}
	}()

	if err := fileWatcher.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func getGtdDir() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln(err)
	}

	notesDir := dirname + "/notes/gtd"

	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.MkdirAll(notesDir, 0755)
	}

	return notesDir
}

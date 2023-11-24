package main

import (
	"bufio"
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

const numFiles = 7

type GtdFile struct {
	name  string
	lines []string
}

func OrganizeGtdFolder(folderPath string) {
	fileData := readGtdFolder(folderPath)
	// output the fileData
	for _, file := range fileData {
		log.Println(file.name)
		for _, line := range file.lines {
			log.Println("    " + line)
		}
	}
}

func readGtdFolder(folderPath string) [numFiles]GtdFile {
	log.Println("Reading your notes:")

	fileNames := [numFiles]string{"in", "next", "done", "future", "projects", "trash", "waiting"}
	fileData := [numFiles]GtdFile{{}, {}, {}, {}, {}, {}, {}}

	for i, file := range fileNames {
		filePath := folderPath + "/" + file + ".norg"
		ensureFileExists(filePath)

		lines := readFile(filePath)
		fileData[i] = GtdFile{file, lines}
	}

	return fileData
}

func ensureFileExists(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

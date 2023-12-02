package main

import (
	"bufio"
	"github.com/radovskyb/watcher"
	"log"
	"os"
	"time"
)

type GtdFile struct {
	Name  string
	Lines []string
}

// See organizeGtdFolder as the "root" of the program
// It does the actual work. Everything else in this file is just based on watching the files
// also look at the parser module

type Tag struct {
	name    string
	tagType TagType
}

type TagType int

const (
	Project = iota
	General
	Who
	Where
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
				organizeGtdFolder(dirname)
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

func organizeGtdFolder(folderPath string) {
	fileData := readGtdFolder(folderPath)
	// output the fileData
	for _, file := range fileData {
		log.Println(file.Name)
		for _, line := range file.Lines {
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

		lines := readFile(filePath)
		fileData[i] = GtdFile{Name: file, Lines: lines}
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

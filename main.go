package main

import (
	"bufio"
	"log"
	"os"

	"github.com/firesquid6/negtd/date"
	"github.com/firesquid6/negtd/gtd"
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
	organizeGtdFolder(dirname)
}

func getGtdDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	notesDir := dirname + "/Dropbox/notes/gtd"

	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		os.MkdirAll(notesDir, 0755)
	}

	return notesDir
}

func organizeGtdFolder(folderPath string) {
	gtdFiles := readGtdFolder(folderPath)
	tasks := []gtd.GtdTask{}
	errors := []string{}
	currentDate := date.GetCurrentDate()

	for _, file := range gtdFiles {
		switch file.Name {
		case "inbox", "backlog":
			newTasks, newErrors := gtd.ReadInboxFile(file.Lines, currentDate)
			tasks = append(tasks, newTasks...)
			errors = append(errors, newErrors...)
		case "agenda":
			newTasks, newErrors := gtd.ReadAgendaFile(file.Lines, currentDate)
			tasks = append(tasks, newTasks...)
			errors = append(errors, newErrors...)
		}
	}
	for _, err := range errors {
		log.Println(err)
	}
	for _, task := range tasks {
		listString := ""
		switch task.GotoList {
		case gtd.Agenda:
			listString = "agenda"
		case gtd.Backlog:
			listString = "backlog"
		case gtd.Inbox:
			listString = "inbox"
		case gtd.Trash:
			listString = "trash"
		}
		log.Println(task.Text + " > " + date.DateToString(task.Date) + " > " + listString)
	}

	inboxFile := gtd.WriteInboxFile(&tasks)
	backlogFile := gtd.WriteBacklogFile(&tasks)
	agendaFile := gtd.WriteAgendaFile(&tasks)

	writeFile(folderPath+"/inbox.norg", inboxFile)
	writeFile(folderPath+"/backlog.norg", backlogFile)
	writeFile(folderPath+"/agenda.norg", agendaFile)
}

const numFiles = 4

func readGtdFolder(folderPath string) [numFiles]GtdFile {
	log.Println("Reading your notes:")

	fileNames := [numFiles]string{"inbox", "agenda", "events", "backlog"}
	fileData := [numFiles]GtdFile{{}, {}, {}, {}}

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

func writeFile(filePath string, content []string) {
	// overwrite the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range content {
		file.WriteString(line + "\n")
	}
}

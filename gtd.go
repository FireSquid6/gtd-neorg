package main

import (
	"bufio"
	"log"
	"os"
)

const numFiles = 7

type GtdFile struct {
	name  string
	lines []string
}

type InTask struct {
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

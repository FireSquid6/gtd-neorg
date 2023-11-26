package main

import (
	"errors"
	"reflect"
	"strings"
)

func parseInboxTask(line string) (GtdTask, error) {
	task := GtdTask{
		text:      "",
		contexts:  []string{},
		project:   "",
		gotoList:  Inbox,
		waitingOn: emptyEvent(),
	}
	if line == "" {
		return task, errors.New("Empty line")
	}

	split := splitBrackets(line)
	split.payload = strings.TrimSpace(split.payload)
	task.text = split.payload

	if split.predata != "" {
		switch split.predata[0] {
		case '!':
			task.gotoList = Next
		case 'x':
			task.gotoList = Done
		case '?':
			task.gotoList = Future
		case '_':
			task.gotoList = Trash
		case '$':
			task.gotoList = Projects
			task.project = split.predata[1:]
		case '%':
			task.gotoList = Waiting
		}
	}

	// parse postdata to get contexts
	if split.postdata != "" {
		split.postdata = strings.TrimSpace(split.postdata)
		split.postdata = strings.ReplaceAll(split.postdata, "@", "")

		contexts := strings.Split(split.postdata, " ")
		for _, context := range contexts {
			task.contexts = append(task.contexts, context)
		}
	}

	return task, nil
}

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}

type bracketSplit struct {
	predata  string
	payload  string
	postdata string
}

func splitBrackets(str string) bracketSplit {
	str = strings.ReplaceAll(str, "[", "|")
	str = strings.ReplaceAll(str, "]", "|")

	hasPre := false
	hasPost := false

	// if the first character is a |, remove it
	if str[0] == '|' {
		hasPre = true
		str = str[1:]
	}

	// if the last character is a |, remove it
	if str[len(str)-1] == '|' {
		hasPost = true
		str = str[:len(str)-1]
	}

	output := bracketSplit{
		predata:  "",
		payload:  "",
		postdata: "",
	}

	// split the string into an array by the | character
	splitStr := strings.Split(str, "|")

	// ugly but works
	if hasPre && hasPost {
		output.predata = splitStr[0]
		output.payload = splitStr[1]
		output.postdata = splitStr[2]
	} else if hasPre {
		output.predata = splitStr[0]
		output.payload = splitStr[1]
	} else if hasPost {
		output.payload = splitStr[0]
		output.postdata = splitStr[1]
	} else {
		output.payload = splitStr[0]
	}

	return output
}

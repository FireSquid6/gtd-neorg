package gtd

import (
	"github.com/firesquid6/negtd/date"
	"reflect"
	"strings"
)

type GtdTask struct {
	Text     string
	GotoList GtdListName
	Date     date.Date
}

type GtdListName int

const (
	Inbox GtdListName = iota
	Agenda
	Backlog
	Trash
	Waiting
)

// INBOX SPECIFIC PARSING
func parseInboxTask(line string, currentDate date.Date) (GtdTask, error) {
	task := GtdTask{
		Text:     "",
		Date:     date.EmptyDate(),
		GotoList: Inbox,
	}

	return task, nil
}

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}

func parseAgendaTask(line string, currentDate date.Date) (GtdTask, error) {
	task := GtdTask{
		Text:     "",
		Date:     date.EmptyDate(),
		GotoList: Agenda,
	}

	return task, nil
}

func trimBeginningWhitespace(input string) string {
	for strings.HasPrefix(input, " ") {
		input = strings.TrimPrefix(input, " ")
	}

	return input
}

func trimEndingWhitespace(input string) string {
	for strings.HasSuffix(input, " ") {
		input = strings.TrimSuffix(input, " ")
	}

	return input
}

func trimWhitespace(input string) string {
	return trimBeginningWhitespace(trimEndingWhitespace(input))
}

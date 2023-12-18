package gtd

import (
	"errors"
	"reflect"
	"strings"

	"github.com/firesquid6/negtd/date"
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
)

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}

func filterPrefix(input string, prefixStart string, prefixEnd string) string {
	inPrefix := false
	prefix := ""
	for _, c := range input {
		if inPrefix {
			if string(c) == prefixEnd {
				inPrefix = false
				break
			}
			prefix += string(c)
		} else {
			if string(c) == prefixStart {
				inPrefix = true
			}
		}
	}

	return prefix
}

func parseInboxTask(line string, currentDate date.Date) (GtdTask, error) {
	task := GtdTask{
		Text:     "",
		Date:     date.EmptyDate(),
		GotoList: Inbox,
	}
	line = trimBeginningWhitespace(line)
	prefix := " "
	if strings.HasPrefix(line, "[") {
		prefix = filterPrefix(line, "[", "]")
		line = strings.TrimPrefix(line, "["+prefix+"] ")
	}
	task.Text = line

	switch string(prefix[0]) {
	case "-", "?":
		task.GotoList = Backlog
	case "_", "x":
		task.GotoList = Trash
	default:
		date, err := date.ParseRelativeDate(prefix, currentDate)
		if err == nil {
			task.Date = date
			task.GotoList = Agenda
		}
	}

	return task, nil
}

func parseAgendaTask(line string, currentDate date.Date) (GtdTask, error) {
	task := GtdTask{
		Text:     "",
		Date:     currentDate,
		GotoList: Agenda,
	}
	line = trimBeginningWhitespace(line)
	if !strings.HasPrefix(line, "- (") {
		return task, errors.New("agenda task must start with '- ('")
	}
	line = strings.TrimPrefix(line, "- ")

	prefix := filterPrefix(line, "(", ")")
	if prefix == "" {
		return task, errors.New("agenda task must have a prefix")
	}
	task.Text = strings.TrimPrefix(line, "("+prefix+") ")
	switch string(prefix[0]) {
	case "-":
		task.GotoList = Backlog
	case "_", "x":
		task.GotoList = Trash
	case ">":
		task.GotoList = Agenda
		dateString := strings.TrimPrefix(prefix, "> ")
		date, err := date.ParseRelativeDate(dateString, currentDate)
		if err != nil {
			return task, err
		}
		task.Date = date
	case "!":
		task.GotoList = Agenda
		task.Date = date.DecrementDate(currentDate)
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

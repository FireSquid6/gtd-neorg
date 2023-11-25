package main

import (
	"errors"
	"testing"
)

func Test_parseInboxTask(t *testing.T) {
	data := []struct {
		name     string
		line     string
		expected GtdTask
		err      error
	}{
		{"empty line", "", GtdTask{}, errors.New("Empty line")},
		{"simple task", "This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Inbox,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task with context", "This is a simple task [@context]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "",
			gotoList:  Inbox,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task with project", "This is a simple task [$project]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "project",
			gotoList:  Inbox,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task with date", "This is a simple task [2020/12/31]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Inbox,
			date:      Date{2020, 12, 31},
			waitingOn: "",
		}, nil},
		{"task with context, project and date", "This is a simple task [@context $project 2020/12/31]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "project",
			gotoList:  Inbox,
			date:      Date{2020, 12, 31},
			waitingOn: "",
		}, nil},
		{"task that needs to be moved to next", "[!] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Next,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task that needs to be moved to done", "[x] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Done,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task that needs to be moved to waiting", "[/] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Waiting,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task that needs to be moved to trash", "[_] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Trash,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task that needs to be moved to future", "[?] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Future,
			date:      Date{-1, -1, -1},
			waitingOn: "",
		}, nil},
		{"task that has something waiting on it", "This is a simple task [\"waiting on some stupid thing\" @context]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "",
			gotoList:  Inbox,
			date:      Date{-1, -1, -1},
			waitingOn: "waiting on some stupid thing",
		}, nil},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			task, err := parseInboxTask(test.line)
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("Expected error %s, got %s", test.err, err)
				}
			} else {
				if !tasksAreEqual(task, &test.expected) {
					t.Errorf("Expected task %v, got %v", test.expected, task)
				}
			}
		})
	}
}

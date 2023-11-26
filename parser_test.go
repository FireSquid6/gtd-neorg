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
			waitingOn: emptyEvent(),
		}, nil},
		{"task with context", "This is a simple task [@context]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "",
			gotoList:  Inbox,
			waitingOn: emptyEvent(),
		}, nil},
		{"task that needs to be moved to next", "[!] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Next,
			waitingOn: emptyEvent(),
		}, nil},
		{"task that needs to be moved to done", "[x] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Done,
			waitingOn: emptyEvent(),
		}, nil},
		{"task that needs to be moved to waiting", "[\"an event\"] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Waiting,
			waitingOn: Event{"an event", emptyDate()},
		}, nil},
		{"task that needs to be moved to trash", "[_] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Trash,
			waitingOn: emptyEvent(),
		}, nil},
		{"task that needs to be moved to future", "[?] This is a simple task", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{},
			project:   "",
			gotoList:  Future,
			waitingOn: emptyEvent(),
		}, nil},
		{"task that has a date to wait for", "[05/15/2021] This is a simple task [@context]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "",
			gotoList:  Inbox,
			waitingOn: Event{"", Date{2021, 5, 15}},
		}, nil},
		{"task that is waiting for a date and has an event", "[05/15/2021 \"an event\"] This is a simple task [@context]", GtdTask{
			text:      "This is a simple task",
			contexts:  []string{"context"},
			project:   "",
			gotoList:  Waiting,
			waitingOn: Event{"an event", Date{2021, 5, 15}},
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

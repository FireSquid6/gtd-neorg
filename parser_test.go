package main

// import (
// 	"errors"
// 	"fmt"
// 	"testing"
// )

// Tests some old code that isn't used anymore. I always forget how to write some for loops so I'm keeping this here
// func Test_parseNextTask(t *testing.T) {
// 	data := []struct {
// 		name     string
// 		line     string
// 		expected GtdTask
// 		err      error
// 	}{
// 		{"empty line", "", GtdTask{}, errors.New("Empty line")},
// 		{"no starting prefix", "I have failed", GtdTask{}, errors.New("No starting prefix")},
// 		{"stay on next", "- ( ) I do nothing and stay", GtdTask{
// 			text:      "I do nothing and stay",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Next,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"throw away", "- ( ) Trash me", GtdTask{
// 			text:      "Trash me",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Trash,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"move to done", "- (x) I am done", GtdTask{
// 			text:      "I am done",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Done,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"move to future", "- (-) send me to the future", GtdTask{
// 			text:      "I am waiting",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Waiting,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 	}
//
// 	fmt.Println(data)
// }
//
// func Test_parseInboxTask(t *testing.T) {
// 	data := []struct {
// 		name     string
// 		line     string
// 		expected GtdTask
// 		err      error
// 	}{
// 		{"empty line", "", GtdTask{}, errors.New("Empty line")},
// 		{"simple task", "This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Inbox,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task with context", "This is a simple task [@context]", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{"context"},
// 			project:   "",
// 			gotoList:  Inbox,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task that needs to be moved to next", "[!] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Next,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task that needs to be moved to done", "[x] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Done,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task that needs to be moved to waiting", "[% \"an event\"] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Waiting,
// 			waitingOn: Event{"an event", emptyDate()},
// 		}, nil},
// 		{"task that needs to be moved to trash", "[_] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Trash,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task that needs to be moved to future", "[?] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "",
// 			gotoList:  Future,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task that has a date to wait for", "[% 05/15/2021] This is a simple task [@context]", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{"context"},
// 			project:   "",
// 			gotoList:  Waiting,
// 			waitingOn: Event{"", Date{5, 15, 2021}},
// 		}, nil},
// 		{"task that is waiting for a date and has an event", "[% 05/15/2021,\"an event\"] This is a simple task [@context]", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{"context"},
// 			project:   "",
// 			gotoList:  Waiting,
// 			waitingOn: Event{"an event", Date{5, 15, 2021}},
// 		}, nil},
// 		{"task that has a project", "[$project] This is a simple task", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{},
// 			project:   "project",
// 			gotoList:  Projects,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 		{"task with multiple contexts", "This is a simple task [@context1 @context2]", GtdTask{
// 			text:      "This is a simple task",
// 			contexts:  []string{"context1", "context2"},
// 			project:   "",
// 			gotoList:  Inbox,
// 			waitingOn: emptyEvent(),
// 		}, nil},
// 	}
//
// 	for _, test := range data {
// 		t.Run(test.name, func(t *testing.T) {
// 			task, err := parseInboxTask(test.line)
// 			if err != nil {
// 				if err.Error() != test.err.Error() {
// 					t.Errorf("Expected error %s, got %s", test.err, err)
// 				}
// 			} else {
// 				if !tasksAreEqual(task, test.expected) {
// 					t.Errorf("Expected task %v, got %v", test.expected, task)
// 				}
// 			}
// 		})
// 	}
// }

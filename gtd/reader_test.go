package gtd

import (
	"github.com/firesquid6/negtd/date"
	"reflect"
	"testing"
)

func Test_ReadInboxFile(t *testing.T) {
	currentDate := date.Date{Year: 2023, Month: 1, Day: 1}

	data := []struct {
		input    []string
		expected []GtdTask
		err      []string
	}{
		{
			input: []string{
				"This is a task",
				"[today] I should go to the agenda",
				"[_] I should be trashed",
				"[?] eh. Send me to the backlog",
			},
			expected: []GtdTask{
				{
					Text:     "This is a task",
					GotoList: Inbox,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "I should go to the agenda",
					GotoList: Agenda,
					Date:     currentDate,
				},
				{
					Text:     "I should be trashed",
					GotoList: Trash,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "eh. Send me to the backlog",
					GotoList: Backlog,
					Date:     date.EmptyDate(),
				},
			},
			err: []string{},
		},
	}

	for _, d := range data {
		tasks, errors := ReadInboxFile(d.input, currentDate)
		if !reflect.DeepEqual(tasks, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, tasks)
		}
		if !reflect.DeepEqual(errors, d.err) {
			t.Errorf("Expected %v, got %v", d.err, errors)
		}
	}

}

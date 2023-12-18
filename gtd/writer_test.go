package gtd

import (
	"reflect"
	"testing"

	"github.com/firesquid6/negtd/date"
)

func Test_WriteInbox(t *testing.T) {
	data := []struct {
		input    []GtdTask
		expected []string
	}{
		{
			input: []GtdTask{
				{
					Text:     "This is a task",
					GotoList: Inbox,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "I should go to the agenda",
					GotoList: Agenda,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "I should be trashed",
					GotoList: Trash,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "this is another task",
					GotoList: Inbox,
					Date:     date.EmptyDate(),
				},
			},
			expected: []string{
				"This is a task",
				"this is another task",
			},
		},
	}

	for _, d := range data {
		tasks := WriteInboxFile(&d.input)
		if !reflect.DeepEqual(tasks, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, tasks)
		}
	}
}

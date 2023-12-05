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
				"(today) I should go to the agenda",
				// "[today] I should go to the agenda [tag1, tag2]", ! this isn't working
				"(_) I should be trashed",
				"(?) eh. Send me to the backlog",
			},
			expected: []GtdTask{
				{
					Text:     "This is a task",
					Tags:     []string{},
					GotoList: Inbox,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "I should go to the agenda",
					Tags:     []string{},
					GotoList: Agenda,
					Date:     currentDate,
				},
				// {
				// 	Text:     "I should go to the agenda",
				// 	Tags:     []string{"tag1", "tag2"},
				// 	GotoList: Agenda,
				// 	Date:     currentDate,
				// },
				{
					Text:     "I should be trashed",
					Tags:     []string{},
					GotoList: Trash,
					Date:     date.EmptyDate(),
				},
				{
					Text:     "eh. Send me to the backlog",
					Tags:     []string{},
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

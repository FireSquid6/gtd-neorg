package gtd

import (
	"reflect"
	"testing"

	"github.com/firesquid6/negtd/date"
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

func Test_ReadAgendaFile(t *testing.T) {
	currentDate := date.Date{Year: 2023, Month: 1, Day: 1}

	data := []struct {
		input    []string
		expected []GtdTask
		err      []string
	}{
		{
			input: []string{
				"* 2023-01-01",
				" - ( ) This is a task",
				"- (today) I should go to the agenda",
				"- (_) I should be trashed",
				"- (-) eh. Send me to the backlog",
				"* 2023-01-02",
				"- ( ) This is a task",
				"- (> today) I should go to the agenda",
				"- (> tomorrow) I am going to tomorrow",
				"- (_) I should be trashed",
				"- (-) eh. Send me to the backlog",
			},
			expected: []GtdTask{
				{
					Text:     "This is a task",
					GotoList: Agenda,
					Date:     currentDate,
				},

				{
					Text:     "I should go to the agenda",
					GotoList: Agenda,
					Date:     currentDate,
				},
				{
					Text:     "I should be trashed",
					GotoList: Trash,
					Date:     currentDate,
				},
				{
					Text:     "eh. Send me to the backlog",
					GotoList: Backlog,
					Date:     currentDate,
				},
				{
					Text:     "This is a task",
					GotoList: Agenda,
					Date:     date.Date{Year: 2023, Month: 1, Day: 2},
				},
				{
					Text:     "I should go to the agenda",
					GotoList: Agenda,
					Date:     date.Date{Year: 2023, Month: 1, Day: 2},
				},
				{
					Text:     "I am going to tomorrow",
					GotoList: Agenda,
					Date: date.Date{
						Year:  2023,
						Month: 1,
						Day:   3,
					},
				},
				{
					Text:     "I should be trashed",
					GotoList: Trash,
					Date:     date.Date{Year: 2023, Month: 1, Day: 2},
				},
				{
					Text:     "eh. Send me to the backlog",
					GotoList: Backlog,
					Date:     date.Date{Year: 2023, Month: 1, Day: 2},
				},
			},
			err: []string{},
		},
	}

	for _, d := range data {
		tasks, errors := ReadAgendaFile(d.input, currentDate)
		for i, task := range tasks {
			if !reflect.DeepEqual(task, d.expected[i]) {
				t.Errorf("Expected %v, got %v", d.expected[i], task)
			}
		}
		if !reflect.DeepEqual(errors, d.err) {
			t.Errorf("Expected %v, got %v", d.err, errors)
		}
	}
}

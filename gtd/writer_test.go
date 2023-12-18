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

func Test_mapTasks(t *testing.T) {
	data := []struct {
		input    []GtdTask
		expected map[date.Date][]GtdTask
	}{
		{
			input: []GtdTask{
				{
					Text:     "This is a task",
					GotoList: Agenda,
					Date:     date.Date{Year: 2023, Month: 1, Day: 1},
				},
				{
					Text:     "I should go to the agenda",
					GotoList: Agenda,
					Date:     date.Date{Year: 2023, Month: 1, Day: 1},
				},
				{
					Text:     "I'm another task",
					GotoList: Agenda,
					Date:     date.Date{Year: 2023, Month: 1, Day: 2},
				},
			},
			expected: map[date.Date][]GtdTask{
				{Year: 2023, Month: 1, Day: 1}: {
					{
						Text:     "This is a task",
						GotoList: Agenda,
						Date:     date.Date{Year: 2023, Month: 1, Day: 1},
					},
					{
						Text:     "I should go to the agenda",
						GotoList: Agenda,
						Date:     date.Date{Year: 2023, Month: 1, Day: 1},
					},
				},
				{Year: 2023, Month: 1, Day: 2}: {
					{
						Text:     "I'm another task",
						GotoList: Agenda,
						Date:     date.Date{Year: 2023, Month: 1, Day: 2},
					},
				},
			},
		},
	}

	for _, d := range data {
		tasks := mapTasks(&d.input)
		if !reflect.DeepEqual(tasks, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, tasks)
		}
	}
}

func Test_WriteAgendaFile(t *testing.T) {
	tasks := []GtdTask{
		{
			Text:     "I'm on the second",
			Date:     date.Date{Year: 2023, Month: 1, Day: 2},
			GotoList: Agenda,
		},
		{
			Text:     "This is a task",
			GotoList: Agenda,
			Date:     date.Date{Year: 2023, Month: 1, Day: 1},
		},
		{
			Text:     "I should go to the agenda",
			Date:     date.Date{Year: 2023, Month: 1, Day: 1},
			GotoList: Agenda,
		},
	}

	expected := []string{
		"* 2023-01-01",
		" - ( ) This is a task",
		" - ( ) I should go to the agenda",
		"",
		"* 2023-01-02",
		" - ( ) I'm on the second",
		"",
	}

	file := WriteAgendaFile(&tasks)
	if len(file) != len(expected) {
		t.Errorf("Expected %v, got %v", len(expected), len(file))
	}

	for i, line := range file {
		if line != expected[i] {
			t.Errorf("Error: %v -> %v", expected[i], line)
		}
	}
}

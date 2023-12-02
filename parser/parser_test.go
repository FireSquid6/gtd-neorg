package parser

import (
	"errors"
	"github.com/firesquid6/negtd/date"
	"reflect"
	"testing"
)

func Test_splitInboxLine(t *testing.T) {
	data := []struct {
		input    string
		expected splitLine
		err      error
	}{
		{
			input: "This is a task",
			expected: splitLine{
				predata:  "",
				text:     "This is a task",
				postdata: "",
			},
			err: nil,
		},
		{
			input: "[between brackets] This is a task",
			expected: splitLine{
				predata:  "between brackets",
				text:     "This is a task",
				postdata: "",
			},
		},
		{
			input: "This has postdata [between brackets]",
			expected: splitLine{
				predata:  "",
				text:     "This has postdata",
				postdata: "between brackets",
			},
		},
	}

	for _, d := range data {
		actual := splitInboxLine(d.input)
		if !reflect.DeepEqual(actual, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, actual)
		}
	}
}

func Test_parseInboxTask(t *testing.T) {
	data := []struct {
		input    string
		expected GtdTask
		err      error
	}{
		{
			input: "This is a task",
			expected: GtdTask{
				text:     "This is a task",
				tags:     []string{},
				gotoList: Inbox,
				date:     date.EmptyDate(),
			},
			err: nil,
		},
		{
			input: "[_] Send me to the trash",
			expected: GtdTask{
				text:     "Send me to the trash",
				tags:     []string{},
				gotoList: Trash,
				date:     date.EmptyDate(),
			},
			err: nil,
		},
		{
			input: "[?] Send me to the backlog",
			expected: GtdTask{
				text:     "Send me to the backlog",
				tags:     []string{},
				gotoList: Backlog,
				date:     date.EmptyDate(),
			},
			err: nil,
		},
		{
			input: "[2023-01-01] Send me to the agenda",
			expected: GtdTask{
				text:     "Send me to the agenda",
				tags:     []string{},
				gotoList: Agenda,
				date: date.Date{
					Year:  2023,
					Month: 1,
					Day:   1,
				},
			},
			err: nil,
		},
		{
			input: "[Today] Send me to the agenda",
			expected: GtdTask{
				text:     "Send me to the agenda",
				tags:     []string{},
				gotoList: Agenda,
				date: date.Date{
					Year:  2023,
					Month: 1,
					Day:   1,
				},
			},
			err: nil,
		}, {
			input: "[Monday] Send me to the agenda",
			expected: GtdTask{
				text:     "Send me to the agenda",
				tags:     []string{},
				gotoList: Agenda,
				date: date.Date{
					Year:  2023,
					Month: 1,
					Day:   2, // January 1st is a Sunday. The next monday is January 2nd.
				},
			},
			err: nil,
		},
		{
			input: "[Sunday] Send me to the agenda",
			expected: GtdTask{
				text:     "Send me to the agenda",
				tags:     []string{},
				gotoList: Agenda,
				date: date.Date{
					Year:  2023,
					Month: 1,
					Day:   8,
				},
			},
			err: nil,
		},
		{
			input: "[Tomorrow] Send me to the agenda",
			expected: GtdTask{
				text:     "Send me to the agenda",
				tags:     []string{},
				gotoList: Agenda,
				date: date.Date{
					Year:  2023,
					Month: 1,
					Day:   2,
				},
			},
			err: nil,
		},
		{
			input: "I have tags [tag1, tag2]",
			expected: GtdTask{
				text:     "I have tags",
				tags:     []string{"tag1", "tag2"},
				gotoList: Inbox,
				date:     date.EmptyDate(),
			},
			err: nil,
		},
	}

	for _, d := range data {
		currentDate := date.Date{
			Year:  2023,
			Month: 1,
			Day:   1,
		}
		actual, err := parseInboxTask(d.input, currentDate)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !tasksAreEqual(actual, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, actual)
		}
	}
}

func Test_splitAgendaTask(t *testing.T) {
	data := []struct {
		input    string
		expected splitLine
		err      error
	}{
		{
			input: "- ( ) This is a task",
			expected: splitLine{
				predata:  "",
				text:     "This is a task",
				postdata: "",
			},
			err: nil,
		},
		{
			input: "- (stuff over here) This is a task",
			expected: splitLine{
				predata:  "stuff over here",
				text:     "This is a task",
				postdata: "",
			},
			err: nil,
		},
		{
			input: "- (stuff over here) This is a task [with postdata]",
			expected: splitLine{
				predata:  "stuff over here",
				text:     "This is a task",
				postdata: "with postdata",
			},
		},
		{
			input: "     - ( ) lots of spaces for no reason",
			expected: splitLine{
				predata:  "",
				text:     "lots of spaces for no reason",
				postdata: "",
			},
		},
		{
			input: "I should fail",
			expected: splitLine{
				predata:  "",
				text:     "",
				postdata: "",
			},
			err: errors.New(""),
		},
	}

	for _, d := range data {
		actual, err := splitAgendaLine(d.input)
		if err != nil && d.err == nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(actual, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, actual)
		}
	}
}

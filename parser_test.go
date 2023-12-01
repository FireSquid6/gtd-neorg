package main

import (
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
				date:     emptyDate(),
			},
			err: nil,
		},
	}

	for _, d := range data {
		actual, err := parseInboxTask(d.input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !tasksAreEqual(actual, d.expected) {
			t.Errorf("Expected %v, got %v", d.expected, actual)
		}
	}
}

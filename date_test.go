package main

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	date, _ := parseDate("2014-01-01")
	if date.year != 2014 || date.month != 1 || date.day != 1 {
		t.Error("Expected 2014-01-01, got ", date)
	}

	date, _ = parseDate("2014-01-1")
	if date.year != 2014 || date.month != 1 || date.day != 1 {
		t.Error("Expected 2014-01-01, got ", date)
	}

	date, _ = parseDate("2023/12/5")
	if date.year != 2023 || date.month != 12 || date.day != 5 {
		t.Error("Expected 2023-12-05, got ", date)
	}

	date, _ = parseDate("2010/6/1")
	if date.year != 2010 || date.month != 6 || date.day != 1 {
		t.Error("Expected 2010-06-01, got ", date)
	}

	date, _ = parseDate("2010/6/01")
	if date.year != 2010 || date.month != 6 || date.day != 1 {
		t.Error("Expected 2010-06-01, got ", date)
	}

	date, err := parseDate("2023/12-6")
	if date.year != 2023 || date.month != 12 || date.day != 6 {
		t.Error("Expected 2023-12-06, got ", date)
	}

	date, err = parseDate("2023/12/6/1")
	if err == nil {
		t.Error("Expected error, got ", date)
	}

	date, err = parseDate("whatever")
	if err == nil {
		t.Error("Expected error, got ", date)
	}

}

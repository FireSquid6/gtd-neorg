package date

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	date, _ := ParseDate("2014-01-01")
	if date.Year != 2014 || date.Month != 1 || date.Day != 1 {
		t.Error("Expected 2014-01-01, got ", date)
	}

	date, _ = ParseDate("2014-01-1")
	if date.Year != 2014 || date.Month != 1 || date.Day != 1 {
		t.Error("Expected 2014-01-01, got ", date)
	}

	date, _ = ParseDate("2023/12/5")
	if date.Year != 2023 || date.Month != 12 || date.Day != 5 {
		t.Error("Expected 2023-12-05, got ", date)
	}

	date, _ = ParseDate("2010/6/1")
	if date.Year != 2010 || date.Month != 6 || date.Day != 1 {
		t.Error("Expected 2010-06-01, got ", date)
	}

	date, _ = ParseDate("2010/6/01")
	if date.Year != 2010 || date.Month != 6 || date.Day != 1 {
		t.Error("Expected 2010-06-01, got ", date)
	}

	date, err := ParseDate("2023/12-6")
	if date.Year != 2023 || date.Month != 12 || date.Day != 6 {
		t.Error("Expected 2023-12-06, got ", date)
	}

	date, err = ParseDate("2023/12/6/1")
	if err == nil {
		t.Error("Expected error, got ", date)
	}

	date, err = ParseDate("whatever")
	if err == nil {
		t.Error("Expected error, got ", date)
	}

}

func Test_getDayOfTheWeek(t *testing.T) {
	date := Date{2023, 11, 27}
	data, err := GetDayOfTheWeek(date)
	if data != 1 {
		t.Error("Expected 3, got ", err)
	}

	date = Date{2023, 12, 6}
	data, err = GetDayOfTheWeek(date)
	if data != 3 {
		t.Error("Expected 4, got", data)
	}

	date = Date{2023, 11, 28}
	data, err = GetDayOfTheWeek(date)
	if data != 2 {
		t.Error("Expected 2, got ", err)
	}

	date = Date{2023, 11, 29}
	data, err = GetDayOfTheWeek(date)
	if data != 3 {
		t.Error("Expected 2, got ", err)
	}
}

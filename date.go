package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func getCurrentDate() Date {
	now := time.Now()
	return Date{now.Year(), int(now.Month()), now.Day()}
}

func dateToString(date Date) string {
	// make sure the month and day are always 2 digits
	month := strconv.Itoa(date.month)
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(date.day)
	if len(day) == 1 {
		day = "0" + day
	}

	return strconv.Itoa(date.year) + "-" + month + "-" + day
}

func parseDate(dateString string) (Date, error) {
	dateString = strings.ReplaceAll(dateString, "-", "/")
	parts := strings.Split(dateString, "/")
	intParts := make([]int, 3)
	date := Date{0, 0, 0}

	if len(parts) == 3 {
		for i, part := range parts {
			intPart, err := strconv.Atoi(part)
			if err != nil {
				return Date{0, 0, 0}, err
			}
			intParts[i] = intPart
		}

		date.day = intParts[2]
		date.month = intParts[1]
		date.year = intParts[0]
	} else {
		return Date{0, 0, 0}, errors.New("Invalid date format. Multiple parts detected")
	}

	return date, nil
}

func datesAreEqual(date1 Date, date2 *Date) bool {
	return date1.year == date2.year && date1.month == date2.month && date1.day == date2.day
}

func parseRelativeDate(dateString Date, currentDate Date) (Date, error) {
	date := Date{0, 0, 0}

	return date, nil
}

func getDayOfTheWeek(date Date) (int, error) {
	dateString := dateToString(date)

	layout := "2006-01-02" // The layout of the input date
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return -1, err
	}

	dayOfWeek := int(t.Weekday())

	return dayOfWeek, nil
}

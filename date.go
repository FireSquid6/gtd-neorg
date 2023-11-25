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

func datesAreEqual(date1 Date, date2 Date) bool {
	return date1.year == date2.year && date1.month == date2.month && date1.day == date2.day
}

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

func parseRelativeDate(dateString string, currentDate Date) (Date, error) {
	lowercaseDateString := strings.ToLower(dateString)
	switch lowercaseDateString {
	case "today":
		return currentDate, nil
	case "tomorrow":
		return incrementDate(currentDate), nil
	case "yesterday":
		return decrementDate(currentDate), nil
	case "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday":
		weekday, err := getWeekdayFromString(lowercaseDateString)
		if err != nil {
			return Date{0, 0, 0}, err
		}
		date, err := getNextDayOfTheWeek(weekday, currentDate)
		if err != nil {
			return Date{0, 0, 0}, err
		}

		return date, nil
	default:
		// try to parse the date
		date, err := parseDate(dateString)
		if err != nil {
			return Date{0, 0, 0}, err
		}
		return date, nil
	}
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

func decrementDate(date Date) Date {
	date.day--

	return validateDate(date)
}

func incrementDate(date Date) Date {
	date.day++

	return validateDate(date)
}

func validateDate(date Date) Date {
	switch date.month {
	case 1, 3, 5, 7, 8, 10, 12:
		if date.day > 31 {
			date.day = 1
			date.month++
		}
	case 4, 6, 9, 11:
		if date.day > 30 {
			date.day = 1
			date.month++
		}
	case 2:
		if date.day > 28 {
			date.day = 1
			date.month++
		}
	}

	if date.month > 12 {
		date.month = 1
		date.year++
	}

	return date
}

func getNextDayOfTheWeek(weekday Weekday, currentDate Date) (Date, error) {
	// get the current day of the week
	currentDayOfWeek, err := getDayOfTheWeek(currentDate)
	if err != nil {
		return Date{0, 0, 0}, err
	}

	// get the difference between the current day of the week and the desired day of the week
	difference := int(weekday) - currentDayOfWeek

	// if the difference is negative or 0, add 7 to it
	if difference <= 0 {
		difference += 7
	}

	// increment the date by the difference
	for i := 0; i < difference; i++ {
		currentDate = incrementDate(currentDate)
	}

	return currentDate, nil
}

func getWeekdayFromString(weekday string) (Weekday, error) {
	switch strings.ToLower(weekday) {
	case "monday":
		return Monday, nil
	case "tuesday":
		return Tuesday, nil
	case "wednesday":
		return Wednesday, nil
	case "thursday":
		return Thursday, nil
	case "friday":
		return Friday, nil
	case "saturday":
		return Saturday, nil
	case "sunday":
		return Sunday, nil
	default:
		return -1, errors.New("Invalid weekday")
	}
}

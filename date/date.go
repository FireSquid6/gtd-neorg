package date

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func EmptyDate() Date {
	return Date{-1, -1, -1}
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func GetCurrentDate() Date {
	now := time.Now()
	return Date{now.Year(), int(now.Month()), now.Day()}
}

func DateToString(date Date) string {
	// make sure the month and day are always 2 digits
	month := strconv.Itoa(date.Month)
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(date.Day)
	if len(day) == 1 {
		day = "0" + day
	}

	return strconv.Itoa(date.Year) + "-" + month + "-" + day
}

func ParseDate(dateString string) (Date, error) {
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

		date.Day = intParts[2]
		date.Month = intParts[1]
		date.Year = intParts[0]
	} else {
		return Date{0, 0, 0}, errors.New("Invalid date format. Multiple parts detected")
	}

	return date, nil
}

func datesAreEqual(date1 Date, date2 *Date) bool {
	return date1.Year == date2.Year && date1.Month == date2.Month && date1.Day == date2.Day
}

func ParseRelativeDate(dateString string, currentDate Date) (Date, error) {
	lowercaseDateString := strings.ToLower(dateString)
	switch lowercaseDateString {
	case "today":
		return currentDate, nil
	case "tomorrow":
		return incrementDate(currentDate), nil
	case "yesterday":
		return decrementDate(currentDate), nil
	case "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday":
		weekday, err := GetWeekdayFromString(lowercaseDateString)
		if err != nil {
			return Date{0, 0, 0}, err
		}
		date, err := GetNextDayOfTheWeek(weekday, currentDate)
		if err != nil {
			return Date{0, 0, 0}, err
		}

		return date, nil
	default:
		// try to parse the date normally
		date, err := ParseDate(dateString)
		if err != nil {
			return Date{0, 0, 0}, err
		}
		return date, nil
	}
}

func GetDayOfTheWeek(date Date) (int, error) {
	dateString := DateToString(date)

	layout := "2006-01-02" // The layout of the input date
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return -1, err
	}

	dayOfWeek := int(t.Weekday())

	return dayOfWeek, nil
}

func decrementDate(date Date) Date {
	date.Day--

	return validateDate(date)
}

func incrementDate(date Date) Date {
	date.Day++

	return validateDate(date)
}

func validateDate(date Date) Date {
	switch date.Month {
	case 1, 3, 5, 7, 8, 10, 12:
		if date.Day > 31 {
			date.Day = 1
			date.Month++
		}
	case 4, 6, 9, 11:
		if date.Day > 30 {
			date.Day = 1
			date.Month++
		}
	case 2:
		if date.Day > 28 {
			date.Day = 1
			date.Month++
		}
	}

	if date.Month > 12 {
		date.Month = 1
		date.Year++
	}

	return date
}

func GetNextDayOfTheWeek(weekday Weekday, currentDate Date) (Date, error) {
	// get the current day of the week
	currentDayOfWeek, err := GetDayOfTheWeek(currentDate)
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

func GetWeekdayFromString(weekday string) (Weekday, error) {
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

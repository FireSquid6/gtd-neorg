package main

import (
	"time"
)

func getCurrentDate() Date {
	now := time.Now()
	return Date{now.Year(), int(now.Month()), now.Day()}
}

func parseDate(dateString string) (Date, error) {

}

func datesAreEqual(date1 Date, date2 Date) bool {
	return date1.year == date2.year && date1.month == date2.month && date1.day == date2.day
}

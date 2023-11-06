package main

import (
	"fmt"
	"time"
)

type Month int

// const (
//	Winter Month = iota + 1
//	Spring
//	Summer
//	Autumn
//	Fall
// )

func main() {
	t := Now()
	fmt.Printf("Today is %s, %s %d, %d\n", t.DayOfWeek, t.Month, t.Day, t.Year)
}

type Time struct {
	Year      int
	Month     string
	Day       int
	DayOfWeek string
}

func Now() Time {
	year, month, day := FromGregorian(time.Now())

	return Time{
		Year:      year,
		Month:     month,
		Day:       day,
		DayOfWeek: DayOfWeek(day),
	}
}

// Constants for seasons and month names
const (
	Winter = "Winter"
	Spring = "Spring"
	Summer = "Summer"
	Autumn = "Autumn"
	Fall   = "Fall"
)

func FromGregorian(date time.Time) (int, string, int) {
	seasons := []string{Winter, Spring, Summer, Autumn, Fall}
	daysInMonth := 73

	// Calculate the total number of days from December 21 of the previous Gregorian year
	dec21 := time.Date(date.Year()-1, time.December, 21, 0, 0, 0, 0, date.Location())
	totalDays := int(date.Sub(dec21).Hours() / 24)

	newCalendarYear := date.Year()
	if totalDays < 11 || totalDays == 365 {
		newCalendarYear++
	}

	// Check if it's a leap day
	if date.Month() == time.February && date.Day() == 29 {
		return newCalendarYear, "Leap Day", 0
	}
	// Check if it's a leap year and after leap day
	if date.Year()%4 == 0 && date.Year()%100 != 0 && totalDays >= 71 {
		totalDays--
	}

	monthIndex := totalDays / daysInMonth
	seasonName := seasons[monthIndex%5]

	// Check if it's a mid-day
	isMidDay := totalDays%daysInMonth == 36
	if isMidDay {
		// Find the corresponding "Mid" day
		return newCalendarYear, fmt.Sprintf("Mid %s", seasonName), 37
	}

	// Determine if it's early or late in the month
	isLateMonth := (totalDays % daysInMonth) >= 36
	if isLateMonth {
		seasonName = "Late " + seasonName
	} else {
		seasonName = "Early " + seasonName
	}
	return newCalendarYear, seasonName, (totalDays % daysInMonth) + 1
}

// Constants for days of the week
const (
	MidSeason = "Mid Season"
	Mercury   = "Mercury"
	Venus     = "Venus"
	Earth     = "Earth"
	Mars      = "Mars"
	Jupiter   = "Jupiter"
	Saturn    = "Saturn"
	Uranus    = "Uranus"
	Neptune   = "Neptune"
	Pluto     = "Pluto"
)

func DayOfWeek(newCalendarDay int) string {
	if newCalendarDay == 37 {
		return MidSeason
	}
	if newCalendarDay == 0 {
		return "Leap Day"
	}
	if newCalendarDay > 37 {
		newCalendarDay--
	}

	// Define the days of the week in the New Calendar
	daysOfWeek := []string{Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto}

	// Calculate the day of the week
	dayIndex := (newCalendarDay - 1) % len(daysOfWeek)
	return daysOfWeek[dayIndex]
}

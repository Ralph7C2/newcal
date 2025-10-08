package newcal

import (
	"fmt"
	"time"
)

// Constants for seasons and month names
const (
	Winter = "Winter"
	Spring = "Spring"
	Summer = "Summer"
	Autumn = "Autumn"
	Fall   = "Fall"
)

var seasons = []string{Winter, Spring, Summer, Autumn, Fall}

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

type Date struct {
	Year  int
	Month string
	Day   int
}

const LeapDay = -1

func (t Date) MonthWithModifier() string {
	modifier := "Early"
	if t.Day > 37 {
		modifier = "Late"
	} else if t.Day == 37 {
		modifier = "Mid"
	}
	return fmt.Sprintf("%s %s", modifier, t.Month)
}

func (t Date) String() string {
	if t.IsLeapDay() {
		return fmt.Sprintf("Leap Day, %d", t.Year)
	}
	dayOfWeek := ""
	if t.Day != 37 {
		dayOfWeek = fmt.Sprintf("%s, ", t.DayOfWeek())
	}
	return fmt.Sprintf("%s%d %s, %d", dayOfWeek, t.Day, t.MonthWithModifier(), t.Year)
}

func (t Date) DayOfWeek() string {
	if t.Day == LeapDay {
		return "Leap Day"
	}
	if t.Day == 37 {
		return MidSeason
	}
	if t.Day > 37 {
		t.Day--
	}

	// Define the days of the week in the New Calendar
	daysOfWeek := []string{Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto}

	// Calculate the day of the week
	dayIndex := (t.Day - 1) % len(daysOfWeek)
	return daysOfWeek[dayIndex]
}

func (t Date) IsLeapDay() bool {
	return t.Day == LeapDay
}

func Parse(date string) (Date, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Date{}, err
	}
	return FromUnix(t.Unix()), nil
}

func FromUnix(unix int64) Date {
	days := unix / 86400
	days += 11
	year := 1970

	for days < 0 {
		if IsLeapYear(year - 1) {
			days += 366
			year--
		} else {
			days += 365
			year--
		}
	}
	for days > 365 {
		if IsLeapYear(year) {
			if days >= 366 {
				days -= 366
				year++
			}
		} else {
			days -= 365
			year++
		}
	}
	if IsLeapYear(year) {
		if days == 70 {
			return Date{
				Year:  year,
				Month: "Leap Day",
				Day:   -1,
			}
		}
		if days > 70 {
			days--
		}
	}

	season := days / 73
	days %= 73

	return Date{
		Year:  year,
		Month: seasons[season],
		Day:   int(days + 1),
	}
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var c = cobra.Command{
	Use:   "newcal [date]",
	Short: "Print a calendar for the given date",
	Run:   run,
}

func main() {
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	t := Now()
	nominal := "Today"
	if len(args) > 0 {
		var err error
		t, err = Parse(args[0])
		if err != nil {
			cmd.PrintErrln(err)
			cmd.Usage()
			return
		}
		nominal = args[0]
	}
	if t.Day == 0 {
		fmt.Printf("%s is LEAP DAY, %d\n", nominal, t.Year)
		return
	}
	if t.Day == 37 {
		fmt.Printf("%s is %s %d, %d\n", nominal, t.Month, t.Day, t.Year)
		return
	}
	fmt.Printf("%s is %s, %s %d, %d\n", nominal, t.DayOfWeek(), t.Month, t.Day, t.Year)

	printCalendar(t)
}

func printCalendar(t Time) {
	fmt.Println("*" + strings.Repeat("-", 79) + "*")
	fmt.Println(formatMonthStr(t.Month))
	fmt.Printf("*-------==-------==-------==-------==-------==-------==-------==-------==-------*\n")
	fmt.Printf("|  Mer  ||  Ven  ||  Ear  ||  Mar  ||  Jup  ||  Sat  ||  Nep  ||  Ura  ||  Plu  |\n")
	fmt.Printf("*-------==-------==-------==-------==-------==-------==-------==-------==-------*\n")
	start, end, weekOffset := 1, 36, 0
	if t.Day > 37 {
		start, end, weekOffset = 38, 73, 1
	}
	for i := start; i <= end; i++ {
		if i == t.Day {
			fmt.Printf("|  *%2d* |", i)
		} else {
			fmt.Printf("| %5d |", i)
		}
		if (i-weekOffset)%9 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("*-------==-------==-------==-------==-------==-------==-------==-------==-------*\n")
}

func formatMonthStr(month string) string {
	monLength := len(month)
	lineLength := 81
	dashes := lineLength - monLength - 2
	leftDashes := dashes / 2
	rightDashes := dashes - leftDashes
	return fmt.Sprintf("|%s%s%s|", strings.Repeat(" ", leftDashes), month, strings.Repeat(" ", rightDashes))
}

type Time struct {
	Year  int
	Month string
	Day   int
}

func Now() Time {
	year, month, day := FromGregorian(time.Now())

	return Time{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func Parse(date string) (Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Time{}, err
	}
	year, month, day := FromGregorian(t)
	return Time{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
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

func (t Time) DayOfWeek() string {
	if t.Day == 37 {
		return MidSeason
	}
	if t.Day == 0 {
		return "Leap Day"
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

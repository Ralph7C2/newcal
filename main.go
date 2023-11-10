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

func fromUnix(unix int64) Time {
	days := unix / 86400
	days += 11
	year := 1970

	for days < 0 {
		if isLeapYear(year - 1) {
			days += 366
			year--
		} else {
			days += 365
			year--
		}
	}
	for days > 365 {
		if isLeapYear(year) {
			if days >= 366 {
				days -= 366
				year++
			}
		} else {
			days -= 365
			year++
		}
	}
	if isLeapYear(year) {
		if days == 70 {
			return Time{
				Year:  year,
				Month: "Leap Day",
				Day:   0,
			}
		}
		if days > 70 {
			days--
		}
	}

	season := days / 73
	days %= 73

	return Time{
		Year:  year,
		Month: seasons[season],
		Day:   int(days + 1),
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func run(cmd *cobra.Command, args []string) {
	t := fromUnix(time.Now().Unix())
	nominal := "Today"
	if len(args) > 0 {
		var err error
		t, err = Parse(args[0])
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Usage()
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

func Parse(date string) (Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Time{}, err
	}
	return fromUnix(t.Unix()), nil
}

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

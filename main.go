package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var c cobra.Command

func init() {
	c = cobra.Command{
		Use:   "newcal",
		Short: "Print a calendar for the given date",
		Run:   run,
	}
	c.Flags().StringP("gregorian", "g", "", "Pass in greogian date to convert(YYYY-MM-DD)")
	c.Flags().StringP("unix", "u", "", "Pass in unix timestamp to convert")
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
	if len(args) > 0 {
		_ = cmd.Usage()
		return
	}

	g, err := cmd.Flags().GetString("gregorian")
	if err != nil {
		cmd.PrintErrln(err)
		_ = cmd.Usage()
		return
	}
	u, err := cmd.Flags().GetString("unix")
	if err != nil {
		cmd.PrintErrln(err)
		_ = cmd.Usage()
		return
	}

	if u != "" && g != "" {
		cmd.PrintErrln("Cannot pass in both unix and gregorian")
		_ = cmd.Usage()
		return
	}

	t := fromUnix(time.Now().Unix())
	nominal := "Today"

	if g != "" {
		var err error
		t, err = Parse(g)
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Usage()
			return
		}
		nominal = g
	}

	if u != "" {
		uu, err := strconv.ParseInt(u, 10, 64)
		if err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Usage()
			return
		}

		t = fromUnix(uu)
		nominal = fmt.Sprintf("%d", uu)
	}

	fmt.Printf("%s is %s\n", nominal, t.String())
	if t.Day == 0 || t.Day == 37 {
		return
	}

	printCalendar(t)
}

func printCalendar(t Time) {
	fmt.Println("*" + strings.Repeat("-", 79) + "*")
	fmt.Println(formatMonthStr(t.MonthWithModifier()))
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

func (t Time) MonthWithModifier() string {
	modifier := "Early"
	if t.Day > 37 {
		modifier = "Late"
	} else if t.Day == 37 {
		modifier = "Mid"
	}
	return fmt.Sprintf("%s %s", modifier, t.Month)
}

func (t Time) String() string {
	if t.Day == 0 {
		return fmt.Sprintf("Leap Day, %d", t.Year)
	}
	dayOfWeek := ""
	if t.Day != 37 {
		dayOfWeek = fmt.Sprintf("%s, ", t.DayOfWeek())
	}
	return fmt.Sprintf("%s%d %s, %d", dayOfWeek, t.Day, t.MonthWithModifier(), t.Year)
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

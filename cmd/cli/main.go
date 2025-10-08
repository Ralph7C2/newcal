package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ralph7c2/newcal"
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

	t := newcal.FromUnix(time.Now().Unix())
	nominal := "Today"

	if g != "" {
		var err error
		t, err = newcal.Parse(g)
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

		t = newcal.FromUnix(uu)
		nominal = fmt.Sprintf("%d", uu)
	}

	fmt.Printf("%s is %s\n", nominal, t.String())
	if t.Day == 0 || t.Day == 37 {
		return
	}

	printCalendar(t)
}

func printCalendar(t newcal.Date) {
	fmt.Println("*" + strings.Repeat("-", 79) + "*")
	fmt.Println(formatMonthStr(t.MonthWithModifier()))
	fmt.Printf("*-------==-------==-------==-------==-------==-------==-------==-------==-------*\n")
	fmt.Printf("|  Mer  ||  Ven  ||  Ear  ||  Mar  ||  Jup  ||  Sat  ||  Ura  ||  Nep  ||  Plu  |\n")
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

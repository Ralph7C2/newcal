package main

import (
	"testing"
	"time"
)

// func TestNow(t *testing.T) {
//	expectedYear := 2023
//	expectedMonth := "Early Fall"
//	expectedDay := 29
//	got := Now()

//	assert.Equal(t, expectedYear, got.Year)
//	assert.Equal(t, expectedMonth, got.Month)
//	assert.Equal(t, expectedDay, got.Day)

//	weekDay := DayOfWeek(got.Day)

//	assert.Equal(t, "Venus", weekDay)
// }

func TestGregorianToNewCalendar(t *testing.T) {
	tests := []struct {
		gregorianDate string
		expectedYear  int
		expectedMonth string
		expectedDay   int
	}{
		{"2022-12-21", 2023, "Early Winter", 1},
		{"2023-01-01", 2023, "Early Winter", 12},
		{"2023-01-02", 2023, "Early Winter", 13},
		{"2023-01-12", 2023, "Early Winter", 23},
		{"2023-01-22", 2023, "Early Winter", 33},
		{"2023-01-25", 2023, "Early Winter", 36},
		{"2023-01-26", 2023, "Mid Winter", 37},
		{"2023-01-31", 2023, "Late Winter", 42},
		{"2023-02-01", 2023, "Late Winter", 43},
		{"2023-02-28", 2023, "Late Winter", 70},
		{"2023-03-03", 2023, "Late Winter", 73},
		{"2023-03-04", 2023, "Early Spring", 1},
		{"2023-03-31", 2023, "Early Spring", 28},
		{"2023-04-01", 2023, "Early Spring", 29},
		{"2023-04-09", 2023, "Mid Spring", 37},
		{"2023-04-10", 2023, "Late Spring", 38},
		{"2023-04-30", 2023, "Late Spring", 58},
		{"2023-05-01", 2023, "Late Spring", 59},
		{"2023-05-15", 2023, "Late Spring", 73},
		{"2023-05-16", 2023, "Early Summer", 1},
		{"2023-05-31", 2023, "Early Summer", 16},
		{"2023-06-01", 2023, "Early Summer", 17},
		{"2023-06-21", 2023, "Mid Summer", 37},
		{"2023-06-22", 2023, "Late Summer", 38},
		{"2023-06-30", 2023, "Late Summer", 46},
		{"2023-07-01", 2023, "Late Summer", 47},
		{"2023-07-27", 2023, "Late Summer", 73},
		{"2023-07-28", 2023, "Early Autumn", 1},
		{"2023-07-31", 2023, "Early Autumn", 4},
		{"2023-08-01", 2023, "Early Autumn", 5},
		{"2023-08-31", 2023, "Early Autumn", 35},
		{"2023-09-01", 2023, "Early Autumn", 36},
		{"2023-09-02", 2023, "Mid Autumn", 37},
		{"2023-09-30", 2023, "Late Autumn", 65},
		{"2023-10-01", 2023, "Late Autumn", 66},
		{"2023-10-08", 2023, "Late Autumn", 73},
		{"2023-10-09", 2023, "Early Fall", 1},
		{"2023-10-31", 2023, "Early Fall", 23},
		{"2023-11-01", 2023, "Early Fall", 24},
		{"2023-11-14", 2023, "Mid Fall", 37},
		{"2023-11-15", 2023, "Late Fall", 38},
		{"2023-11-30", 2023, "Late Fall", 53},
		{"2023-12-01", 2023, "Late Fall", 54},
		{"2023-12-20", 2023, "Late Fall", 73},
		{"2023-12-21", 2024, "Early Winter", 1},
		{"2024-01-01", 2024, "Early Winter", 12},
		{"2024-02-28", 2024, "Late Winter", 70},
		{"2024-02-29", 2024, "Leap Day", 0},
		{"2024-03-01", 2024, "Late Winter", 71},
		{"2024-03-03", 2024, "Late Winter", 73},
	}

	for _, test := range tests {
		gregorianDate, _ := time.Parse("2006-01-02", test.gregorianDate)
		year, month, day := FromGregorian(gregorianDate)

		if year != test.expectedYear || month != test.expectedMonth || day != test.expectedDay {
			t.Errorf("For Gregorian Date %s, expected: (%d, %s, %d), got: (%d, %s, %d)", test.gregorianDate, test.expectedYear, test.expectedMonth, test.expectedDay, year, month, day)
		}
	}
}

func TestDayOfWeek(t *testing.T) {
	tests := []struct {
		newCalDay int
		expected  string
	}{
		{0, "Leap Day"},
		{1, "Mercury"},
		{2, "Venus"},
		{3, "Earth"},
		{4, "Mars"},
		{5, "Jupiter"},
		{6, "Saturn"},
		{7, "Uranus"},
		{8, "Neptune"},
		{9, "Pluto"},
		{10, "Mercury"},
		{11, "Venus"},
		{29, "Venus"},
		{36, "Pluto"},
		{37, "Mid Season"},
		{38, "Mercury"},
	}

	for _, test := range tests {
		got := DayOfWeek(test.newCalDay)
		if got != test.expected {
			t.Errorf("For New Calendar Day %d, expected: %s, got: %s", test.newCalDay, test.expected, got)
		}
	}
}

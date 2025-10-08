package newcal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromUnix(t *testing.T) {
	assert.Equal(t, Date{Year: 1968, Month: "Autumn", Day: 73}, FromUnix(time.Date(1968, 10, 8, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1968, Month: "Fall", Day: 1}, FromUnix(time.Date(1968, 10, 9, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1968, Month: "Fall", Day: 73}, FromUnix(time.Date(1968, 12, 20, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1969, Month: "Winter", Day: 1}, FromUnix(time.Date(1968, 12, 21, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1969, Month: "Fall", Day: 73}, FromUnix(time.Date(1969, 12, 20, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 1}, FromUnix(time.Date(1969, 12, 21, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 11}, FromUnix(-86400))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 11}, FromUnix(time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 12}, FromUnix(0))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 12}, FromUnix(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 13}, FromUnix(86400))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 13}, FromUnix(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 71}, FromUnix(5097600))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 71}, FromUnix(time.Date(1970, 3, 1, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Winter", Day: 73}, FromUnix(time.Date(1970, 3, 3, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1970, Month: "Spring", Day: 1}, FromUnix(time.Date(1970, 3, 4, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1971, Month: "Winter", Day: 71}, FromUnix(36633600))
	assert.Equal(t, Date{Year: 1971, Month: "Winter", Day: 71}, FromUnix(time.Date(1971, 3, 1, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1972, Month: "Leap Day", Day: 0}, FromUnix(36633600+365*86400))
	assert.Equal(t, Date{Year: 1972, Month: "Leap Day", Day: 0}, FromUnix(time.Date(1972, 2, 29, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1972, Month: "Winter", Day: 71}, FromUnix(36633600+366*86400))
	assert.Equal(t, Date{Year: 1987, Month: "Spring", Day: 73}, FromUnix(time.Date(1987, 5, 15, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1999, Month: "Summer", Day: 1}, FromUnix(time.Date(1999, 5, 16, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 2999, Month: "Summer", Day: 73}, FromUnix(time.Date(2999, 7, 27, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 3147, Month: "Autumn", Day: 1}, FromUnix(time.Date(3147, 7, 28, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 3147, Month: "Autumn", Day: 73}, FromUnix(time.Date(3147, 10, 8, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 1, Month: "Autumn", Day: 73}, FromUnix(time.Date(1, 10, 8, 0, 0, 0, 0, time.UTC).Unix()))
	assert.Equal(t, Date{Year: 0, Month: "Autumn", Day: 73}, FromUnix(time.Date(0, 10, 8, 0, 0, 0, 0, time.UTC).Unix()))
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
		got := Date{Day: test.newCalDay}.DayOfWeek()
		if got != test.expected {
			t.Errorf("For New Calendar Day %d, expected: %s, got: %s", test.newCalDay, test.expected, got)
		}
	}
}

package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	pDate, _ := time.Parse("January 2, 2006 15:04:05 ", date)
	return time.Now().After(pDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	pDate, _ := time.Parse(layout, date)
	pHour := pDate.Hour()
	return pHour >= 12 && pHour < 18
}

//IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	cDate, _ := time.Parse(layout, date)
	formattedString := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", cDate.Weekday(), cDate.Month(), cDate.Day(), cDate.Year(), cDate.Hour(), cDate.Minute())
	return string(formattedString)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

// Mon Jan 2 15:04:05 MST 2006

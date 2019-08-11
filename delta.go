package delta

import (
	"math"
	"time"
)

const (
	startHour = 10
	endHour   = 18
)

func adjustDay(t time.Time) time.Time {
	var adjust time.Duration
	if twd := t.Weekday(); twd == time.Sunday {
		adjust = 24
	} else if twd == time.Saturday {
		adjust = 48
	}
	if adjust == 0 {
		return t
	}
	m := t.Add(time.Hour * adjust)
	return time.Date(m.Year(), m.Month(), m.Day(), startHour, 0, 0, 0, time.Local)
}

func adjustHour(t time.Time) time.Time {
	if t.Hour() < startHour {
		return time.Date(t.Year(), t.Month(), t.Day(), startHour, 0, 0, 0, time.Local)
	}
	if t.Hour() > endHour {
		m := t.Add(time.Hour * 24)
		return time.Date(m.Year(), m.Month(), m.Day(), startHour, 0, 0, 0, time.Local)
	}
	return t
}

func Delta(start, end time.Time) (float64, int, int) {
	// Adjust start to be Monday if in the weekend
	start = start.In(time.Local)
	start = adjustHour(start)
	start = adjustDay(start)
	end = end.In(time.Local)
	end = adjustHour(end)
	end = adjustDay(end)

	delta := end.Sub(start).Hours()
	days := math.Round(delta / 24)
	weekends := math.Trunc(days / 7)
	if days > 0 && end.Weekday() < start.Weekday() {
		weekends += 1
	}
	return end.Sub(start).Hours() - 16*(days+weekends), int(days), int(weekends)
}

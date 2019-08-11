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

func removeHolidays(start, end time.Time) float64 {
	var adjust float64
	yearStart, yearEnd := start.Year(), end.Year()
	for year := yearStart; year <= yearEnd; year++ {
		holidays := []time.Time{
			time.Date(year, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(year, 5, 8, 0, 0, 0, 0, time.Local),
			time.Date(year, 7, 14, 0, 0, 0, 0, time.Local),
			time.Date(year, 8, 15, 0, 0, 0, 0, time.Local),
			time.Date(year, 11, 11, 0, 0, 0, 0, time.Local),
			time.Date(year, 12, 25, 0, 0, 0, 0, time.Local),
		}
		for _, holiday := range holidays {
			if wd := holiday.Weekday(); wd == time.Saturday || wd == time.Sunday {
				continue
			}
			if start.Before(holiday) && end.After(holiday) {
				adjust += 8
			}
		}
	}
	return adjust
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
	return end.Sub(start).Hours() - 16*(days+weekends) - removeHolidays(start, end), int(days), int(weekends)
}

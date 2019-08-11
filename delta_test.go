package delta

import (
	"testing"
	"time"
)

func TestDelta(t *testing.T) {
	data := []struct {
		a           time.Time
		b           time.Time
		want        float64
		wantDay     int
		wantWeekend int
		name        string
	}{
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 5, 14, 0, 0, 0, time.Local),
			want:        2,
			wantDay:     0,
			wantWeekend: 0,
			name:        "same day",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 5, 13, 30, 0, 0, time.Local),
			want:        1.5,
			wantDay:     0,
			wantWeekend: 0,
			name:        "same day half an hour",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 14, 0, 0, 0, time.Local),
			want:        10,
			wantDay:     1,
			wantWeekend: 0,
			name:        "overnight",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 11, 0, 0, 0, time.Local),
			want:        7,
			wantDay:     1,
			wantWeekend: 0,
			name:        "overnight earlier",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 2, 0, 0, 0, time.Local),
			want:        6,
			wantDay:     1,
			wantWeekend: 0,
			name:        "overnight early",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 20, 0, 0, 0, time.Local),
			want:        14,
			wantDay:     2,
			wantWeekend: 0,
			name:        "overnight late",
		},
		{
			a:           time.Date(2019, time.August, 5, 15, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 10, 0, 0, 0, time.Local),
			want:        3,
			wantDay:     1,
			wantWeekend: 0,
			name:        "overnight start hour",
		},
		{
			a:           time.Date(2019, time.August, 8, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 13, 14, 0, 0, 0, time.Local),
			want:        26,
			wantDay:     5,
			wantWeekend: 1,
			name:        "overweek",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 20, 14, 0, 0, 0, time.Local),
			want:        82,
			wantDay:     14,
			wantWeekend: 2,
			name:        "overmultiweek",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 20, 10, 0, 0, 0, time.Local),
			want:        78,
			wantDay:     14,
			wantWeekend: 2,
			name:        "overmultiweek ealier",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 19, 14, 0, 0, 0, time.Local),
			want:        74,
			wantDay:     13,
			wantWeekend: 2,
			name:        "overmultiweek ealier day",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 21, 14, 0, 0, 0, time.Local),
			want:        90,
			wantDay:     15,
			wantWeekend: 2,
			name:        "overmultiweek later day",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 21, 20, 0, 0, 0, time.Local),
			want:        94,
			wantDay:     16,
			wantWeekend: 2,
			name:        "overmultiweek later day late night",
		},
		{
			a:           time.Date(2019, time.August, 6, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 21, 2, 0, 0, 0, time.Local),
			want:        86,
			wantDay:     15,
			wantWeekend: 2,
			name:        "overmultiweek later day early",
		},
		{
			a:           time.Date(2019, time.August, 5, 0, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 14, 0, 0, 0, time.Local),
			want:        12,
			wantDay:     1,
			wantWeekend: 0,
			name:        "early",
		},
		{
			a:           time.Date(2019, time.August, 5, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 6, 20, 0, 0, 0, time.Local),
			want:        14,
			wantDay:     2,
			wantWeekend: 0,
			name:        "late",
		},
		{
			a:           time.Date(2019, time.August, 2, 14, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 2, 20, 0, 0, 0, time.Local),
			want:        4,
			wantDay:     3,
			wantWeekend: 1,
			name:        "late friday",
		},
		{
			a:           time.Date(2019, time.August, 3, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 5, 11, 0, 0, 0, time.Local),
			want:        1,
			wantDay:     0,
			wantWeekend: 0,
			name:        "start saturday",
		},
		{
			a:           time.Date(2019, time.August, 4, 12, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 5, 11, 0, 0, 0, time.Local),
			want:        1,
			wantDay:     0,
			wantWeekend: 0,
			name:        "start sunday",
		},
		{
			a:           time.Date(2019, time.August, 2, 14, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 3, 12, 0, 0, 0, time.Local),
			want:        4,
			wantDay:     3,
			wantWeekend: 1,
			name:        "end saturday",
		},
		{
			a:           time.Date(2019, time.August, 2, 14, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 4, 12, 0, 0, 0, time.Local),
			want:        4,
			wantDay:     3,
			wantWeekend: 1,
			name:        "end sunday",
		},
		{
			a:           time.Date(2019, time.August, 3, 14, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 4, 12, 0, 0, 0, time.Local),
			want:        0,
			wantDay:     0,
			wantWeekend: 0,
			name:        "start end weekend",
		},
		{
			a:           time.Date(2019, time.August, 2, 4, 0, 0, 0, time.Local),
			b:           time.Date(2019, time.August, 4, 12, 0, 0, 0, time.Local),
			want:        8,
			wantDay:     3,
			wantWeekend: 1,
			name:        "start early end sunday",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			got, days, weekends := Delta(d.a, d.b)
			if got != d.want || days != d.wantDay || weekends != d.wantWeekend {
				t.Errorf("Delta(%v, %v): got %v, %v days, %v weekends; want %v, %v days, %v weekends", d.a, d.b, got, days, weekends, d.want, d.wantDay, d.wantWeekend)
			}
		})
	}
}

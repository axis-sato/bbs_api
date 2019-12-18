package pkg

import "time"

var fakeTime time.Time

func Freeze(t time.Time) {
	fakeTime = t
}

func Reset() {
	fakeTime = time.Time{}
}

func Now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

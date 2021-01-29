package clock

import "time"

var fakeTime time.Time

func SetFakeTime(t time.Time) {
	fakeTime = t
}

func ResetFake() {
	fakeTime = time.Time{}
}

func Now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

func getJst() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

func JstNow() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}

	return time.Now().UTC().In(getJst())
}

func GetMonthFirst(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, getJst())
}

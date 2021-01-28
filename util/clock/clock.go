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

func JstNow() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Now().UTC().In(jst)
}

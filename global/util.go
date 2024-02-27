package global

import "time"

var (
	TimeLayout = "2006-01-02 15:04:05"
)

func NowTimeToString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TodayMidnight() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
}

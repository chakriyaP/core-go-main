package funcDatetime

import (
	"time"
)


func DateNow() string {
	time_location, _ := time.LoadLocation("Asia/Bangkok")
	time_current := time.Now().In(time_location)
	time_format := time_current.Format("2006-01-02") //ตั้งค่า Format
	return time_format
}

func TimeNow() string {
	time_location, _ := time.LoadLocation("Asia/Bangkok")
	time_current := time.Now().In(time_location)
	time_format := time_current.Format("15:04:05") //ตั้งค่า Format
	return time_format
}

func TimeDelaySec(n time.Duration) {
	time.Sleep(n * time.Second)
}

func TimeDelayMinute(n time.Duration) {
	time.Sleep(n * time.Minute)
	/// time.Millisecond
}

func TimeDelayMillisecond(n time.Duration) {
	time.Sleep(n * time.Millisecond)
}
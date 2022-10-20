package date

import (
	"fmt"
	"time"
)

func Sleep(num int64) {
	time.Sleep(time.Duration(num) * time.Second)
}

var loadLocation = "Local"

// 时间戳
func GetTime() int64 {
	return (time.Now().Unix())
}

// 当前时间
func GetDate() string {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetDateStr() string {
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetDateDay() string {
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
}

func GetDateYMD() string {
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
}

func GetDateToTime(dateStr string) int64 {
	loc, _ := time.LoadLocation(loadLocation)
	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	unix_time := the_time.Unix()
	return unix_time
}

func GetTimeToDate(times int64) string {
	unix_time := time.Unix(times, 0)
	return unix_time.Format("2006-01-02 15:04:05")
}

func GetTodayStartAndEndTime() (int64, int64) {
	loc, _ := time.LoadLocation(loadLocation)
	date := time.Now().Format("2006-01-02")
	startDate := date + " 00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate, loc)
	endDate := date + " 23:59:59"
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate, loc)
	return startTime.Unix(), endTime.Unix()
}

func GetTodayStartAndEndDate() (string, string) {
	startTime, endTime := GetTodayStartAndEndTime()
	return GetTimeToDate(startTime), GetTimeToDate(endTime)
}

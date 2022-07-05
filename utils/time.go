package utils

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

const (
	TimeLayout         string = "2006-01-02 15:04:05"
	TimeLayoutNum      string = "20060102150405"
	TimeLayoutDay      string = "2006-01-02"
	TimeLayoutDayNum   string = "20060102"
	TimeLayoutMonth    string = "2006-01"
	TimeLayoutMonthNum string = "200601"
)

var UtlTime = &_time{}

type _time struct {
}

func (o *_time) TimeToDateNum(t time.Time) string {
	return t.Format(TimeLayoutNum)
}

func (o *_time) UnixToDate(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format(TimeLayout)
}

func (o *_time) UnixToDateNum(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayoutNum)
}

func (o *_time) UnixToDay(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayoutDay)
}

func (o *_time) UnixToDayNum(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayoutDayNum)
}

func (o *_time) UnixToMonth(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayoutMonth)
}

func (o *_time) UnixToMonthNum(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayoutMonthNum)
}

func (o *_time) DateToUnix(date string) int64 {
	var layout string
	switch len(date) {
	case 10:
		layout = TimeLayoutDay
	case 19:
		layout = TimeLayout
	default:
		log.Println(fmt.Sprintf("[DateToTimestampLenError] dateStr: [%s] dateLen: %d", date, len(date)))
		return 0
	}

	t, e := time.ParseInLocation(layout, date, time.Local)
	if e != nil {
		log.Println(fmt.Sprintf("[DateToTimestampError] dateStr: [%s] err: %s", date, e.Error()))
		return 0
	}
	return t.Unix()
}

func (o *_time) ValidateDate(dayStr string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return reg.MatchString(dayStr)
}

func (o *_time) ValidateMinute(minuteStr string) bool {
	reg := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !reg.MatchString(minuteStr) {
		return false
	}
	if minuteStr[:2] > "23" {
		return false
	}
	if minuteStr[3:] > "59" {
		return false
	}
	return true
}

func (o *_time) ValidateTime(timeStr string) bool {
	reg := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`)
	if !reg.MatchString(timeStr) {
		return false
	}
	if timeStr[:2] > "23" {
		return false
	}
	if timeStr[3:5] > "59" {
		return false
	}
	if timeStr[6:] > "59" {
		return false
	}
	return true
}

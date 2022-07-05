package utils

import (
	"fmt"
	"strings"
	"time"
)

// 时间戳转年月日 时分秒
func DateFormat(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04")
}

// 时间戳转年月日 时分秒
func DateFormatYmdHis(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 时间戳转年月日 时分秒
func DateFormatYmdHis64(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 时间戳转时分秒
func DateFormatHis(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("15:04:05")
}

// 时间戳转时分秒
func DateFormatHi(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("15:04")
}

//时间戳转年月日
func DateFormatYmd(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02")
}

//获取当前年月
func DateYmFormat() string {
	tm := time.Now()
	return tm.Format("2006-01")
}

//获取当前年月-汉字
func DateHanYmFormat(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006年01月")
}

//获取年月日时分秒（字符串类型）
func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

//获取年月日时分秒（字符串类型rfc3339） second = +60 or -60
func DateNowRfc3339Str(second string) string {
	m, _ := time.ParseDuration(second + "s")
	tm := time.Now().Add(m)
	return tm.Format("2006-01-02T15:04:05+08:00")
	// 2018-06-08T10:34:56+08:00
}

//rfc3339时间转时间戳（字符串类型rfc3339） strTime := "2018-03-24T20:01:00+08:00"
func DateRfc3339Unix(strTime string) int64 {
	t, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", strTime, time.Local)
	if err != nil {
		fmt.Println("err", err)
	}
	return t.Unix()
}

//获取年月日时分秒（字符串类型遵循rfc3339标准格式）
func DateFormatYmdhis() string {
	tm := time.Now()
	ymdhis := tm.Format("20060102150405")
	return ymdhis
}

//获取年月日（字符串类型遵循rfc3339标准格式）
func DateYmd() string {
	tm := time.Now()
	ymd := tm.Format("20060102")
	return ymd
}

//获取年月（字符串类型遵循rfc3339标准格式）
func DateYm() string {
	tm := time.Now()
	ym := tm.Format("200601")
	return ym
}

//获取当前年月
func DateYmFormatHi() string {
	tm := time.Now()
	return tm.Format("15:04")
}

//获取当前时分秒
func DateYmFormatHis() string {
	tm := time.Now()
	return tm.Format("15:04:05")
}

//时间戳
func DateUnix() int {
	t := time.Now().Local().Unix()
	return int(t)
}
func DateUnix64() int64 {
	t := time.Now().Local().Unix()
	return t
}

//获取年月日时分秒(time类型)
func DateNowFormat() time.Time {
	tm := time.Now()
	return tm
}

//获取第几周
func DateWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}

//获取年、月、日
func DateYMD() (int, int, int) {
	year, month, day := DateYmdInts()
	return year, month, day
}

// 获取年月日
func DateYmdFormat() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

// 获取日期的年月日
func DateYmdInts() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}

//获取当前时间的微秒数
func DateGetMicroTime() int64 {
	return time.Now().UnixNano() / 1000000
}

//日期获取时间戳
func DateGetUnix(date string) int64 {
	//获取本地location
	//toBeCharge := "2015-01-01 23:59:59"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                       //转化所需模板
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                      //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	return sr
}

//获取年月日的时间戳
func DateGetUnixYmd(date string) int64 {
	//获取本地location
	//toBeCharge := "2015-01-01 23:59:59"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02"                                //转化所需模板
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                      //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	return sr
}

//获取年月日的时间戳 20060102
func DateGetYmd(date string) int64 {
	//获取本地location
	//toBeCharge := "2015-01-01 23:59:59"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "20060102"                                  //转化所需模板
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                      //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	return sr
}

//RFC3339 时间格式化
func RFC3339ToCSTLayout(value string) string {

	timeLayout := "2006-01-02 15:04:05"
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return ""
	}
	loc, _ := time.LoadLocation("Local")

	return ts.In(loc).Format(timeLayout)
}

/**
获取本周周一的日期
*/
func DateGetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取本周周一的日期后几天的日期
*/
func DateGetFirstDateOfWeekDay(d int) (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day()+d, 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取本周周日的日期
*/
func DateGetLastDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	offset += 6
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取上周的周一日期
*/
func DateGetLastWeekFirstDate() (weekMonday string) {
	thisWeekMonday := DateGetFirstDateOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format("2006-01-02")
	return
}

//获取本月的第一天和最后一天时间戳
func DateGetMonthFirstLastDay() (int64, int64) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Unix(), lastOfMonth.Unix() + 60*60*24 - 1
}

//获取本月的第一天和最后一天时间戳
func DateGetMonthFirstLastDayByAddMonth(addMonth int) (int64, int64) {
	now := time.Now().AddDate(0, addMonth, 0)
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Unix(), lastOfMonth.Unix() + 60*60*24 - 1
}

//获取星期几
func DateGetWeekDay() int {
	t := time.Now()
	return int(t.Weekday())
}

//获取今日起始时间戳
func DateGetStartEndUnix() (int64, int64) {
	currDayStartUnix := DateGetUnix(DateYmdFormat() + " 00:00:00")
	currDayEndUnix := DateGetUnix(DateYmdFormat() + " 23:59:59")
	return currDayStartUnix, currDayEndUnix
}

//获取当前剩余时间，单位秒
func DateGetDaySurplus() int64 {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	return todayLastTime.Unix() - time.Now().Local().Unix()
}

//获取一个月以后的时间
func DateGetAfterMonthDate(timeStr string) (afterMonth string) {
	TimeStr, _ := time.Parse("20060102", timeStr)
	afterMonthTime := TimeStr.AddDate(0, 1, 0)
	afterMonth = afterMonthTime.Format("20060102")
	return
}

//获取昨天的年月日2006-01-02 格式
func DateYestDayTime() (yesterday string) {
	timeStr := DateYmdFormat()
	TimeStr, _ := time.Parse("2006-01-02", timeStr)
	yesterdayTime := TimeStr.AddDate(0, 0, -1)
	yesterday = yesterdayTime.Format("2006-01-02")
	return
}

//获取昨天的年月日20060102 格式
func DateYestDayTimeFormt() (yesterday string) {
	timeStr := DateYmd()
	TimeStr, _ := time.Parse("20060102", timeStr)
	yesterdayTime := TimeStr.AddDate(0, 0, -1)
	yesterday = yesterdayTime.Format("20060102")
	return
}

//获取昨天起始结束时间戳
func DateYestDayTimeUnix() (int64, int64) {
	dayStr := DateYestDayTime()
	currDayStartUnix := DateGetUnix(dayStr + " 00:00:00")
	currDayEndUnix := DateGetUnix(dayStr + " 23:59:59")
	return currDayStartUnix, currDayEndUnix
}

//获取上月月份 2006-01，1
func DateBeforeMonth(YMStr string, num int) string {
	TimeStr, _ := time.Parse("2006-01", YMStr)
	beforMonthTime := TimeStr.AddDate(0, -num, 0)
	return beforMonthTime.Format("2006-01")
}

/**
 * 星期数字 转 汉字
 * @param $w
 * @return string
 */
func DateWeekChinse(w int) string {
	wChinse := ""
	switch w {
	case 0:
		wChinse = "日"
	case 1:
		wChinse = "一"
	case 2:
		wChinse = "二"
	case 3:
		wChinse = "三"
	case 4:
		wChinse = "四"
	case 5:
		wChinse = "五"
	case 6:
		wChinse = "六"
	case 7:
		wChinse = "日"
	default:
		wChinse = "未知"
	}
	return wChinse
}

/**
 * 获取星期数组 转为汉字
 * @param $weeks
 * @return string
 */
func DateWeeksChinse(weeks string) string {
	limitWeekArr := Split(weeks, ",")
	weeksCHinse := ""
	for _, o := range limitWeekArr {
		if weeksCHinse != "" {
			weeksCHinse += ","
		}
		weeksCHinse += DateWeekChinse(StringToInt(o))
	}
	return weeksCHinse
}

//根据时分秒格式化yyyy-MM-dd HH:mm:ss
func DateGetyMdHmsByhms(hms string) string {
	hmsArr := strings.Split(hms, ":")
	h := "00"
	m := "00"
	s := "00"
	if len(hmsArr) > 0 {
		h = hmsArr[0]
		if len(h) == 1 {
			h = "0" + h
		}
	}
	if len(hmsArr) > 1 {
		m = hmsArr[1]
		if len(m) == 1 {
			m = "0" + m
		}
	}
	if len(hmsArr) > 2 {
		s = hmsArr[2]
		if len(s) == 1 {
			s = "0" + s
		}
	}
	Ymd := DateYmdFormat()
	return fmt.Sprintf("%v %v:%v:%v", Ymd, h, m, s)
}

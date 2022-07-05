package holidays

import (
	"errors"
	"github.com/txbao/common-go/utils"
	"time"
)

//获取指定日期 T + n 日期
func GetTNDay(value string, n int) (date string, err error) {
	for i := 1; i <= n; i++ {
		date, err = GetTOneDay(value)
		//fmt.Println("date-:", date)
		if err != nil {
			return
		}
		value = date
	}
	return
}

//获取T+n时间戳 his为时间如 04:15:12
func GetTNUnix(his string, n int) (int64, error) {
	currDay := utils.Date("2006/01/02", time.Now().Local().Unix())
	date, err := GetTNDay(currDay, n)
	if err != nil {
		return 0, err
	}
	dateFormat := utils.StrReplace("/", "-", date, -1)
	if his == "" {
		his = utils.DateYmFormatHis()
	}
	currYmdHis := dateFormat + " " + his
	return utils.DateGetUnix(currYmdHis), nil
}

//获取指定日期的T+1 日期
func GetTOneDay(value string) (string, error) {
	collectionTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return "", errors.New("日期错误")
	}
	//隔天
	nextDay := collectionTime.Add(24 * time.Hour)
	if verifyDate(nextDay) {
		return nextDay.Format("2006/01/02"), nil
	}
	return GetTOneDay(nextDay.Format("2006/01/02"))
}

//获取T+1时间戳 his为时间如 04:15:12
func GetT1Unix(his string) (int64, error) {
	currDay := utils.Date("2006/01/02", time.Now().Local().Unix())
	date, err := GetTOneDay(currDay)
	if err != nil {
		return 0, err
	}
	dateFormat := utils.StrReplace("/", "-", date, -1)
	if his == "" {
		his = utils.DateYmFormatHis()
	}
	currYmdHis := dateFormat + " " + his
	return utils.DateGetUnix(currYmdHis), nil
}

func verifyDate(value time.Time) bool {
	//是否为周末
	weekInt := int(value.Weekday())
	if weekInt == 6 || weekInt == 0 {
		return false
	}
	//是否为节假日
	res := IsHoliday(value.Format("2006/01/02"))
	if res {
		return false
	}
	return true
}

func IsWorkDay(value string) bool {
	if IsHoliday(value) {
		return false
	}
	collectionTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return false
	}
	isWorkDay := int(collectionTime.Weekday())
	if isWorkDay == 0 || isWorkDay == 6 {
		return false
	}
	return true
}

//返回是否节假日
func IsHoliday(value string) (result bool) {
	collectionTime, err := time.Parse("2006/01/02", value)
	if err != nil {
		return
	}
	nowYear, _, _ := time.Now().Date()
	if collectionTime.Year() > nowYear+1 {
		return
	}
	collections := FetchByYear(collectionTime.Year())
	for _, collection := range collections {
		startDate, _ := getDate(collection.Start)
		endDate, _ := getDate(collection.End)
		if collectionTime.Unix() >= startDate.Unix() && collectionTime.Unix() <= endDate.Unix() {
			return true
		}
	}
	return

}

// FetchByYear get holidays by year in china
func FetchByYear(year int) []OneCollection {
	var index int
	nowYear, _, _ := time.Now().Date()
	if year > nowYear+1 {
		return nil
	}
	index = nowYear - year
	//fmt.Println(index)
	if index < len(FetchCollectionYearHistory().Data) && index >= 0 {
		return FetchCollectionYearHistory().Data[index]
	}
	return []OneCollection{}
}

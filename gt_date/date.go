package gt_date

import "time"

// GetWeekDateRange 获取指定日期所在周的日期范围
// @param date time.Time 指定日期
// @return startDate time.Time 周的开始日期
// @return endDate time.Time 周的结束日期
func GetWeekDateRange(date time.Time) (startDate, endDate time.Time) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	weekDay := int(date.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	startDate = date.AddDate(0, 0, -weekDay+1)
	endDate = startDate.AddDate(0, 0, 6)
	return
}

// GetCurrWeekDateRange 获取当前周的日期范围
// @return startDate time.Time 周的开始日期
// @return endDate time.Time 周的结束日期
func GetCurrWeekDateRange() (startDate, endDate time.Time) {
	return GetWeekDateRange(time.Now())
}

// GetLastWeekDateRange 获取上周的日期范围
// @return startDate time.Time 周的开始日期
// @return endDate time.Time 周的结束日期
func GetLastWeekDateRange() (startDate, endDate time.Time) {
	now := time.Now()
	last := now.AddDate(0, 0, -7)
	return GetWeekDateRange(last)
}

package util

import "time"

const TimeConst = "15:04:05"
const DateConst = "2006-01-01"
const DateTimeConst = DateConst + " " + TimeConst

func Time2Str(t time.Time) string {
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(TimeConst)
	return str
}

func Str2Time(formatTimeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeConst, formatTimeStr, loc)
	return theTime
}

func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

func Time2Stamp(t time.Time) int64 {
	millisecond := t.UnixNano() / 1e6
	return millisecond
}

func Stamp2Str(stamp int64) string {
	str := time.Unix(stamp/1000, 0).Format(TimeConst)
	return str
}

func Stamp2Time(stamp int64) time.Time {
	stampStr := Stamp2Str(stamp)
	timer := Str2Time(stampStr)
	return timer
}

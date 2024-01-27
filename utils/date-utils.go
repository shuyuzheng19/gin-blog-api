package utils

import "time"

func TimeStampToTime(millTimeStamp int64) time.Time {
	return time.UnixMilli(millTimeStamp)
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

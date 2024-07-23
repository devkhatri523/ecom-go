package utils

import (
	"strconv"
	"time"
)

func GetUTCTime() time.Time {
	return time.Now().UTC()
}
func GetTime() time.Time {
	return time.Now()
}
func ParseRFC3339Time(tm string) time.Time {
	parseTime, err := time.Parse(time.RFC3339, tm)
	if err != nil {
		panic(err)
	}
	return parseTime
}
func Timestamp() int64 {
	now := time.Now()
	return now.UnixMilli()
}
func parseTs(ts string) time.Time {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0).UTC()
}
func ParseTimestamp(ts string) int64 {
	return parseTs(ts).UnixMilli()
}
func FormatStrTimestamp(ts string) string {
	return FormatTime(parseTs(ts))
}
func FormatTimestamp(ts int64) string {
	return FormatTime(parseTs(strconv.FormatInt(ts, 10)))
}
func TimestampToTime(timestamp int64) time.Time {
	return time.UnixMilli(timestamp)
}

func TimestampToTimeWithLayout(timestamp int64, layout string) string {
	return time.UnixMilli(timestamp).Format(layout)
}
func FormatTimestampWithLayout(ts int64, layout string) string {
	return parseTs(strconv.FormatInt(ts, 10)).Format(layout)
}
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}
func GetUTCDateTime() string {
	now := time.Now().UTC()
	return FormatTime(now)
}

func GetUTCMySqlTime() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02 03:04:05")
}

func ParseUTCMySqlTime(tm string) time.Time {
	parseTime, err := time.Parse("2006-01-02 03:04:05", tm)
	if err != nil {
		panic(err)
	}
	return parseTime
}

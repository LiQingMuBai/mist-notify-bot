package tools

import (
	"fmt"
	"time"
)

// 常用时间格式常量
const (
	FormatDate            = "2006-01-02"                    // 年月日
	FormatTime            = "15:04:05"                      // 时分秒
	FormatDateTime        = "2006-01-02 15:04:05"           // 年月日 时分秒
	FormatDateTimeNoSpace = "2006-01-02_15:04:05"           // 年月日_时分秒(文件命名友好)
	FormatDateTimeCompact = "20060102150405"                // 年月日时分秒(紧凑格式)
	FormatDateTimeMilli   = "2006-01-02 15:04:05.000"       // 带毫秒
	FormatDateTimeMicro   = "2006-01-02 15:04:05.000000"    // 带微秒
	FormatDateTimeNano    = "2006-01-02 15:04:05.000000000" // 带纳秒
	FormatDateTimeChinese = "2006年01月02日 15时04分05秒"         // 中文格式
	FormatDateTime12Hour  = "2006-01-02 03:04:05 PM"        // 12小时制
)

func Format4Chinesese(t time.Time) string {
	return t.Format(FormatDateTimeChinese)
}

// FormatNow 格式化当前时间
func FormatNow(layout string) string {
	return time.Now().Format(layout)
}

// Format 格式化指定时间
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatYMD 格式化为年月日(2006-01-02)
func FormatYMD(t time.Time) string {
	return t.Format(FormatDate)
}

// FormatHMS 格式化为时分秒(15:04:05)
func FormatHMS(t time.Time) string {
	return t.Format(FormatTime)
}

// FormatYMDHMS 格式化为年月日 时分秒(2006-01-02 15:04:05)
func FormatYMDHMS(t time.Time) string {
	return t.Format(FormatDateTime)
}

// FormatCompact 格式化为紧凑格式(20060102150405)
func FormatCompact(t time.Time) string {
	return t.Format(FormatDateTimeCompact)
}

// FormatForFilename 格式化为文件名友好格式(2006-01-02_15-04-05)
func FormatForFilename(t time.Time) string {
	return t.Format(FormatDateTimeNoSpace)
}

// FormatWithMilli 格式化为带毫秒的时间
func FormatWithMilli(t time.Time) string {
	return t.Format(FormatDateTimeMilli)
}

// FormatChinese 格式化为中文时间格式
func FormatChinese(t time.Time) string {
	return t.Format(FormatDateTimeChinese)
}

// Format12Hour 格式化为12小时制时间
func Format12Hour(t time.Time) string {
	return t.Format(FormatDateTime12Hour)
}

// FormatByTimestamp 将时间戳格式化为字符串
func FormatByTimestamp(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// FormatByTimestampNano 将纳秒时间戳格式化为字符串
func FormatByTimestampNano(timestamp int64, layout string) string {
	return time.Unix(0, timestamp).Format(layout)
}

// FormatDuration 格式化时间间隔
func FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

package providers

import "time"

// parseTime 解析时间字符串
func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

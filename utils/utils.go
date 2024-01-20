package utils

// StringIfNull 返回第一个不为空的字符串
func StringIfNull(str ...string) string {
	for _, item := range str {
		if item != "" {
			return item
		}
	}
	return ""
}

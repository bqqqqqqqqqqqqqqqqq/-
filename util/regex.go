package util

import "regexp"

// 邮箱正则表达式
func IsEmail(str ...string) bool {
	if str == nil {
		return false
	}
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", s)
		if false == b {
			return b
		}
	}
	return b
}

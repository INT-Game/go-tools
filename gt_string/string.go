package gt_string

import "unicode/utf8"

// SubStrDecodeRuneInString 截取字符串的前 length 个字符
func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}
	return s[:n]
}

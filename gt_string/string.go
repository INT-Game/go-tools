package gt_string

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// SubStrDecodeRuneInString 截取字符串的前 length 个字符
func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}
	return s[:n]
}

func ConvertToString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

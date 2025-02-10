package gt_sign

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5String 获取字符串的MD5值
// @param bytes []byte 字节数组
// @return md5Str string MD5值
func GetMd5String(bytes []byte) (md5Str string) {
	h := md5.New()
	h.Write(bytes)
	md5Str = hex.EncodeToString(h.Sum(nil))
	return
}

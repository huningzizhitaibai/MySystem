package common

func TranslateForMe(src string) string {
	utf8Bytes := []byte(src)

	// 将字节切片转换为字符串
	answer := string(utf8Bytes)
	return answer
}

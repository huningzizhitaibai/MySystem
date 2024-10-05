package common

import (
	"log"
	"strconv"
)

func ReplaceForMe(src string) string {
	unescapedStr, err := strconv.Unquote(`"` + src + `"`)
	if err != nil {
		log.Fatalf("Failed to unquote string: %v", err)
	}
	return unescapedStr
	// 打印转换后的字符串
	//fmt.Println(unescapedStr) // 输出: 你好

	//    // 查找 'b' 字符的位置
	//index := strings.IndexByte(temp, 'b')
	//if index == -1 {
	//	// 没有找到 'b'，返回原始字符串
	//	return temp
	//}
	//// 返回删除第一个 'b' 后的字符串
	//return temp[:index] + temp[index+1:]

}

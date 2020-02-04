package LeetCode

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/yezihack/algo/00.src"
	"math"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"
)
//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
/*
示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN (−231) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MyAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	//判断是否是空格
	str = strings.ReplaceAll(str, " ", "")
	//判断第一个元素
	first := str[0]
	if !unicode.IsDigit(rune(first)) {
		return 0
	}
	//正则提取数字与-符号
	reg := regexp.MustCompile(`^\d+|\-|\.`)
	ss := reg.FindAllString(str, -1)
	fmt.Println(ss)

	//切片转换成字符串
	ans := ""
	for _, s := range ss {
		if s == "." {
			break
		}
		ans += s
	}
	//转换成整型
	fmt.Println(ans)
	ani, err := strconv.ParseInt(ans, 10, 64)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	for ani < math.MinInt32 {
		return math.MinInt32
	}
	if  ani > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(ani)
}
func TestMyAtoi(t *testing.T) {
	src.Asset(1, t, -42, MyAtoi("  -42"))
	src.Asset(2, t, 4193, MyAtoi("4193 with words"))
	src.Asset(3, t, 0, MyAtoi("words and 987"))
	src.Asset(4, t, -2147483648, MyAtoi("-91283472332"))
	src.Asset(5, t, 3, MyAtoi("3.14159"))
}
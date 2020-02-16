package _2_数学

import (
	"fmt"
	"testing"
)

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//解题思路: 先建立hash,映射罗马与数字之间的关系,需要处理一个特殊的关系,
//如IV, I表示1, V表示5, 整个表示4, 即V-I=4.所以我们需要想到后项与前项比较.
//如果后大于前,则减, 如果小于则加. 如IV, 先减, 得负数-1, 然后再加5,最终结果就是4.

*/
//执行用时 :
//28 ms
//, 在所有 Go 提交中击败了
//10.15%
//的用户
//内存消耗 :
//4.2 MB
//, 在所有 Go 提交中击败了
//20.79%
//的用户
func RomanToInt(s string) int {
	hash := make(map[string]int)
	hash["I"] = 1
	hash["V"] = 5
	hash["X"] = 10
	hash["L"] = 50
	hash["C"] = 100
	hash["D"] = 500
	hash["M"] = 1000
	ans := 0
	for i := 0; i < len(s); i ++ {
		if i < len(s)-1 && hash[string(s[i])] < hash[string(s[i+1])] {
			ans -= hash[string(s[i])]
		} else {
			ans += hash[string(s[i])]
		}
		fmt.Printf("%s,ans:%d,,i:%d,s[%d]:%c\n", s,ans, i, i, s[i])
	}
	return ans
}
func TestRomanToInt(t *testing.T) {
	tests := []struct {
		index int
		data string
		expect int
	}{
		{1, "III", 3},
		{2, "IV", 4},
		{3, "IX", 9},
		{4, "LVIII", 58},
		{5, "MCMXCIV", 1994},
	}
	for _, item := range tests {
		if ret := RomanToInt(item.data); ret != item.expect {
			t.Errorf("expect:%v, actual:%v\n", item.expect, ret)
		}
	}
}
//执行用时 :
//8 ms
//, 在所有 Go 提交中击败了
//81.28%
//的用户
//内存消耗 :
//3.1 MB
//, 在所有 Go 提交中击败了
//93.19%
//的用户
func RomanToIntV2(s string) int {
	sum := 0
	getValue := func(char byte) int {
		switch char {
		case 'I':return 1
		case 'V':return 5
		case 'X':return 10
		case 'L':return 50
		case 'C':return 100
		case 'D':return 500
		case 'M':return 1000
		default:return 0
		}
	}
	//获取第一个值.相当于前值
	preNum := getValue(s[0])
	for i := 1;i < len(s); i ++{
		//获取第二个值. 相当于后值
		num := getValue(s[i])
		//前值与后值比较, 如果后值大于前值,说明需要减去前值.
		if preNum < num {
			sum -= preNum
		} else {//否则正常加上.
			sum += preNum
		}
		//继续保持前后值的关系.
		preNum = num
	}
	//最后将后值加上.
	sum += preNum
	return sum
}
func TestRomanToIntV2(t *testing.T) {
	tests := []struct {
		index int
		data string
		expect int
	}{
		//{1, "III", 3},
		//{2, "IV", 4},
		{3, "IX", 9},
		//{4, "LVIII", 58},
		//{5, "MCMXCIV", 1994},
	}
	for _, item := range tests {
		if ret := RomanToIntV2(item.data); ret != item.expect {
			t.Errorf("expect:%v, actual:%v\n", item.expect, ret)
		}
	}
}
func RomanToIntV3(s string) int {
	hash := make(map[byte]int)
	hash['I'] = 1
	hash['V'] = 5
	hash['X'] = 10
	hash['L'] = 50
	hash['C'] = 100
	hash['D'] = 500
	hash['M'] = 1000
	sum := 0
	//前项的值与后项的值进行比较, 如果后大于前,则减去, 因为IV表示4, 则I=1, V=5, 即V-I, 5-1=4
	//防止溢出,则len(s)-1, 还有一位没有加上,则最后处理
	for i := 0; i < len(s)-1; i ++ {
		if hash[s[i]] < hash[s[i+1]] {
			sum -= hash[s[i]]
		} else {
			sum += hash[s[i]]
		}
		fmt.Printf("%s,ans:%d,,i:%d,s[%d]:%c\n", s, sum, i, i, s[i])
	}
	//处理剩余的字符
	sum += hash[s[len(s)-1]]
	return sum
}
func TestRomanToIntV3(t *testing.T) {
	tests := []struct {
		index int
		data string
		expect int
	}{
		{1, "III", 3},
		{2, "IV", 4},
		{3, "IX", 9},
		{4, "LVIII", 58},
		{5, "MCMXCIV", 1994},
	}
	for _, item := range tests {
		if ret := RomanToIntV3(item.data); ret != item.expect {
			t.Errorf("expect:%v, actual:%v\n", item.expect, ret)
		}
	}
}
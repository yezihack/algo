package other_test

import (
	"fmt"
	"testing"
)

//使用临时变量交互值
func TestCrossValueTemp(t *testing.T) {
	a, b := 10, 20
	fmt.Printf("pre: a:%d, b:%d\n", a, b)
	tmp := a
	a = b
	b = tmp
	fmt.Printf("later: a:%d, b:%d\n", a, b)
}

//不使用临时变量交互值
//思路: 首先a = a + b , 然后b = a - b 相当于,把a = a + b代入进来就是这样啦 b = (a + b) - b
//然后b = a - b 相当于 a = (a + b) - a (为什么是a 因为在第二步b 已经是 a的值啦), a相互抵消. a = b
func TestCrossValue(t *testing.T) {
	a, b := 10, 20
	fmt.Printf("pre: a:%d, b:%d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("later: a:%d, b:%d\n", a, b)
}

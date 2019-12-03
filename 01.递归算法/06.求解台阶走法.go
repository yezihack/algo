package Recurive

//求解台阶走法. 1个台阶的走法是1种方式, 2个台阶的走法是2种方式.
func Step(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Step(n - 1) + Step(n -  2)
}
var cache = make(map[int]int)
func StepMap(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := cache[n]; ok {
		return val
	}
	ret := StepMap(n-1) + StepMap(n-2)
	cache[ret] = ret
	return ret
}
func StepF(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	pre := 2
	prepre := 1
	for i := 3; i <= n; i ++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}
	return ret
}

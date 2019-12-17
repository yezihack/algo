package string_string

func LeftShiftOne(s string) string {
	b := []byte(s)
	one := b[0]
	for i := 1;i < len(b); i ++ {
		b[i-1] = b[i]
	}
	b[len(b)-1] = one
	return string(b)
}

func LeftShiftTwo(s string) string {
	b := []byte(s)
	t1 := b[0]
	t2 := b[1]
	for i := 2;i < len(b);i ++ {
		b[i-2] = b[i]
	}
	b[len(b)-1] = t2
	b[len(b)-2] = t1
	return string(b)
}

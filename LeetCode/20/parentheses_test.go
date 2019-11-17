package _0

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestIsValid(t *testing.T) {
	src.So(t, true, IsValid("(([]))"))
	src.So(t, false, IsValid("([)]"))
	src.So(t, true, IsValid("()[]{}"))
	src.So(t, false, IsValid("(]"))
	src.So(t, false, IsValid(")"))
}
func TestIsValidBySlice(t *testing.T) {
	//src.So(t, true, IsValidBySlice("()"))
	//src.So(t, true, IsValidBySlice("({[]})"))
	//src.So(t, false, IsValidBySlice("({[{]})"))
	//src.So(t, false, IsValidBySlice("("))
	//src.So(t, false, IsValidBySlice("(}"))
	//src.So(t, false, IsValidBySlice("(}()"))
	src.So(t, false, IsValidBySlice("]"))
}
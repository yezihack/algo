package string_string

import (
	"fmt"
	"testing"
)

func TestLeftShiftOne(t *testing.T) {
	fmt.Println(LeftShiftOne(LeftShiftOne("abcdef")))
}
func TestLeftShiftTwo(t *testing.T) {
	//abcdef->cdefab
	fmt.Println(LeftShiftTwo("abcdef"))

}
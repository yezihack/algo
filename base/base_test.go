package base

import (
	"fmt"
	"math"
	"testing"
)

func TestLog(t *testing.T) {
	fmt.Println(math.Log2(8))
	fmt.Println(math.Log2(128))
	fmt.Println(math.Log2(128 * 2))
	fmt.Println(math.Log2(1000000000))
}

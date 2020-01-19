package LeetCode

import(
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestPow2(t *testing.T) {
	src.Asset(1, t, float64(1024), Pow2(2, 10))
}

func TestPow3(t *testing.T) {
	src.Asset(1, t, float64(1024), Pow3(2, 10))
}
func TestPow4(t *testing.T) {
	src.Asset(1, t, float64(1024), Pow4(2, 10))
	src.Asset(1, t, float64(2048), Pow4(2, 11))
}
func TestPow5(t *testing.T) {
	src.Asset(1, t, float64(1024), Pow5(2, 10))
}

//169
func TestMajorityElement(t *testing.T) {
	src.Asset(1, t, 3, MajorityElement([]int{3,2,3}))
}
func TestMajorityElement2(t *testing.T) {
	src.Asset(1, t, 3, MajorityElement2([]int{3,2,3}))
	src.Asset(2, t, 3, MajorityElement2([]int{1, 1, 4, 5,2,3}))
	src.Asset(3, t, 3, MajorityElement2([]int{0, 0, 1, 1, 1, 56, 8, 9, 10, 22}))
}
//122
func TestMaxProfit(t *testing.T) {
	src.Asset(1, t, 7, MaxProfit([]int{7,1,5,3,6,4}))
	src.Asset(2, t, 4, MaxProfit([]int{1,2,3,4,5}))
	src.Asset(3, t, 0, MaxProfit([]int{7,6,4,3,1}))
}
func TestMaxProfit2(t *testing.T) {
	src.Asset(1, t, 5, MaxProfit2([]int{7,1,5,3,6,4}))
}
func TestMaxProfit3(t *testing.T) {
	src.Asset(1, t, 5, MaxProfit3([]int{7,1,5,3,6,4}))
}

func TestTT(t *testing.T) {
	fmt.Println(1, 1 & 1, 1 % 2)
	fmt.Println(2, 2 & 1, 2 % 2)
	fmt.Println(3, 3 & 1, 3 % 2)
	fmt.Println(4, 4 & 1, 4 % 2)
	fmt.Println(5, 5 & 1, 5 % 2)
}

func TestUniqueOccurrences(t *testing.T) {
	src.Asset(1, t, true, UniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	src.Asset(2, t, false, UniqueOccurrences([]int{1, 2}))
	src.Asset(3, t, true, UniqueOccurrences([]int{-3,0,1,-3,1,1,1,-3,10,0}))
	src.Asset(4, t, false, UniqueOccurrences([]int{26,2,16,16,5,5,26,2,5,20,20,5,2,20,2,2,20,2,16,20,16,17,16,2,16,20,26,16}))
}
func TestUniqueOccurrences2(t *testing.T) {
	src.Asset(4, t, false, UniqueOccurrences2([]int{26,2,16,16,5,5,26,2,5,20,20,5,2,20,2,2,20,2,16,20,16,17,16,2,16,20,26,16}))
}
package _3_动态规划

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestPubStr(t *testing.T) {
	src.Asset(1, t, "lue", PubStr("clues","blue"))
	src.Asset(2, t, "fsh", PubStr("fish","fosh"))
	src.Asset(3, t, "llo", PubStr("hello","kallo"))
	src.Asset(4, t, "ks", PubStr("kiss","kas"))
}

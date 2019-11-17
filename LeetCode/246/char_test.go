package _46

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	src.Asset(1, t, true, IsAnagram("cat", "tac"))
	src.Asset(2, t, true, IsAnagram("caaat", "tacaa"))
	src.Asset(3, t, false, IsAnagram("family", "familp"))
}
func TestIsAnagramByMap(t *testing.T) {
	src.Asset(1, t, true, IsAnagramByMap("qwe", "ewq"))
	src.Asset(2, t, false, IsAnagramByMap("fast", "tsag"))
	src.Asset(3, t, true, IsAnagramByMap("s", "s"))
	src.Asset(4, t, false, IsAnagramByMap("pq", "qq"))
}
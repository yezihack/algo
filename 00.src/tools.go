package src

import (
	"fmt"
	"testing"
)

//断言函数, 简单快捷好用
func So(t *testing.T, expect, actual interface{}) {
	if expect != actual {
		t.Errorf("%s, expect:%v, actual:%v\n", t.Name(), expect, actual)
	}
}
//断言函数,带name
func Asset(name interface{}, t *testing.T, expect, actual interface{}) {
	var nameStr string
	switch name.(type) {
	case int:
		nameStr = fmt.Sprint(name)
	case string:
		nameStr = fmt.Sprint(name)
	default:
		nameStr = "default"
	}
	t.Run(nameStr, func(t *testing.T) {
		if expect != actual {
			t.Errorf("%s, expect:%v, actual:%v\n", t.Name(), expect, actual)
		}
	})
}

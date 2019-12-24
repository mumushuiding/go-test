package extend

import "testing"

func Test1(t *testing.T) { // 子类调用父类方法
	apple := &apple{}
	apple.types = "apple"
	apple.name = "a1"
	apple.SayTypes()
}

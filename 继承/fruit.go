package extend

import "fmt"

type fruit struct {
	types string
}

func (f *fruit) SayTypes() {
	fmt.Println("START----------------")
	fmt.Println(f) // 基本无法打印，派生类自有属性
	fmt.Printf("成功调用基类方法 SayTypes: %s\n", f.types)
	fmt.Println("END----------------")
}

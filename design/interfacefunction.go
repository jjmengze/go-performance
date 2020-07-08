package main

import (
	"fmt"
)

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

// Each 我不把話說死，只要有實作  Handler interface 的我可以以接受，反正我是呼叫 Handler interface 的 Do function
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// EachFunc 總之你就是傳入一個 type HandlerFunc 型態的誤進給我就對
// 基本上傳入的 function 會實作 Handler interface，所以我在裡面又讓他傳入 Each。
func EachFunc(m map[interface{}]interface{}, fun HandlerFunc) {
	Each(m, fun)
}

func selfInfo(k, v interface{}) {
	fmt.Printf("我叫%s,今年%d歲\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["A同學"] = 20
	persons["小明"] = 23
	persons["老王"] = 26


	//我認為這樣設計可以把要用的function 直接暴露出來，有利使用者直接閱讀或是操作。
	EachFunc(persons, selfInfo)

	//另外透過定義interface的方式可以讓使用者選擇傳入有時做該interface的struct，會更加的彈性與靈活。
	Each(persons, HandlerFunc(selfInfo))

}

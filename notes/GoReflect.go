package notes

import (
	"fmt"
	"reflect"
)

func GoReflectNotes() {

	// reflect 实现了运行时的反射能力，能够让程序操作不通类型的对象。
	// 反射包中有两对非常重要的函数和类型，分别是
	// reflect.TypeOf  能获取类型信息
	// reflect.ValueOf  能获取数据的运行时表示

	// reflect.Type 是反射包定义的一个接口
	// 我们可以使用reflect.TypeOf函数获取任意变量的类型
	a := 1
	fmt.Println(reflect.TypeOf(a))

	// reflect.Value的类型与 reflect.Type不同，它被声明成了结构体。
	// 这个结构体没有对外暴露的字段，但是提供了获取或者写入数据的方法。

	// 反射包中的所有方法基本都是围绕着 reflect.Type 和 reflect.Value 两个类型设计的。
	// 我们通过 reflect.TypeOf、reflect.ValueOf
	// 可以将一个普通的变量转换成反射包中提供的 reflect.Type 和 reflect.Value
	// 随后就可以使用反射包中的方法对它们进行复杂的操作。

	// 三大法则
	// 从interface{}变量可以反射出反射对象
	// 从反射对象可以获取interface{}变量
	// 要修改反射对象，其值必须可以设置

}

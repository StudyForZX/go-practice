// assert的详细文档
// https://pkg.go.dev/github.com/stretchr/testify@v1.9.0/assert#Assertions

package testify

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {

	// 简写方式，否则下边方法都得传t参数
	assert := assert.New(t)

	// Equal断言
	assert.Equal(4, 4, "The result should be 4")
	assert.NotEqual(4, 3, "The result should be 4")

	// 判断零值
	assert.Zero(0, "The number should be 0")
	assert.NotZero(1, "The number should not be 0")

	// 当使用Exactly时，需要类型跟值都一样
	// assert.Exactly(int32(123), int64(123))
}

func TestString(t *testing.T) {
	// 可以判断字符串包含
	assert.Contains(t, "Hello World", "World")
	assert.NotContains(t, "Hello World", "go")
}

func TestBoolean(t *testing.T) {
	assert.True(t, 1+1 == 2, "1+1==2 should be true")
	assert.False(t, 1+1 == 3, "1+1==3 should be false")
	// 1s内，闭包函数内的值都为false，每10毫秒检查一次
	assert.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
}

func TestArray(t *testing.T) {
	// 可以判断切片包含
	assert.Contains(t, [2]string{"Hello", "World"}, "World")
}

func TestSlice(t *testing.T) {
	// 可以判断切片包含
	assert.Contains(t, []string{"Hello", "World"}, "World")

	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	// equal判断引用类型时，仅判断值是否相同
	// 注意：这里的切片要保障顺序是一样的，否则不相等
	assert.Equal(t, sl1, sl2, "sl1 should equal to sl2")

	sl3 := []int{1, 3, 2}
	// 这里的切片要保障顺序是一样的，否则不相等
	assert.NotEqual(t, sl1, sl3, "sl1 should not equal to sl2")

	// 不论顺序判断每个值是否相同，如果有重复的需要对比重复的值是否相等
	assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 2, 3})
}

func TestMap(t *testing.T) {
	// 可以判断map包含，仅能查key
	assert.Contains(t, map[string]string{"Hello": "World"}, "Hello")
	assert.NotContains(t, map[string]string{"Hello": "World"}, "World")

	map1 := map[string]string{"m1": "1", "m2": "2"}
	map2 := map[string]string{"m2": "2", "m1": "1"}
	assert.Equal(t, map1, map2)
}

func TestPointer(t *testing.T) {
	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	sl3 := []int{2, 3, 4}
	p1 := &sl1
	p2 := &sl2
	// equal判断指针类型时，判断的是指针对应的值是否相同
	assert.Equal(t, p1, p2, "The content which p1 point to should equal to which p2 point to")
	assert.NotEqual(t, sl1, sl3, "sl1 should not equal to sl3")
}

type TestA struct {
	A1 int `json:"a1"`
	A2 int `json:"a2"`
}

func TestError(t *testing.T) {
	err := errors.New("demo error")
	assert.EqualError(t, err, "demo error")

	t1 := &TestA{A1: 1, A2: 2}
	t1Json, err := json.Marshal(t1)
	assert.Zero(t, err, string(t1Json))
}

type TestB struct {
	B1 *TestB1
}

type TestB1 struct {
	TB11 string
	TB12 int
	TB13 *TestB2
}

type TestB2 struct {
	TB21 int
	TB22 bool
}

func TestStruct(t *testing.T) {
	b1 := &TestB{
		B1: &TestB1{
			TB11: "11",
			TB12: 12,
			TB13: nil,
		},
	}

	p1 := &TestB1{
		TB11: "11",
		TB12: 12,
	}

	assert.Equal(t, b1.B1, p1)
}

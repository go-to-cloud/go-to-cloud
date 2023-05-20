package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJsonTime_MarshalJSON(t *testing.T) {
	// 创建一个 JsonTime 类型的实例
	jt := JsonTime(time.Date(2022, 4, 9, 22, 30, 0, 0, time.UTC))

	// 执行 MarshalJSON 方法，将 JsonTime 转换为 JSON 字符串
	b, err := jt.MarshalJSON()
	assert.NoError(t, err)

	// 将 JSON 字符串反序列化
	jt2 := JsonTime{}
	err = jt2.UnmarshalJSON(b)
	assert.NoError(t, err)

	// 比较反序列化后的 time.Time 是否与原始 JsonTime 相等
	t1 := time.Time(jt)
	t2 := time.Time(jt2)
	assert.Equal(t, t1, t2)
}

func TestJsonDate_MarshalJSON(t *testing.T) {
	// 创建一个 JsonDate 类型的实例
	jd1 := JsonDate(time.Date(2022, 4, 9, 0, 0, 0, 0, time.UTC))

	// 执行 MarshalJSON 方法，将 JsonDate 转换为 JSON 字符串
	b, err := jd1.MarshalJSON()
	assert.NoError(t, err)

	// 将 JSON 字符串反序列化为 time.Time 类型
	jd2 := JsonDate{}
	err = jd2.UnmarshalJSON(b)

	d1 := time.Time(jd1)
	d2 := time.Time(jd2)
	assert.Equal(t, d1, d2)
}

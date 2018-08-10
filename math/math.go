package math

import (
	"encoding/json"

	simplejson "github.com/bitly/go-simplejson"
)

// Math math struct
type Math struct {
	A int
	B int
}

// New new instance
func New(a, b int) *Math {
	return &Math{
		A: a,
		B: b,
	}
}

// Divide 除法
func (m *Math) Divide() int {
	// FIX ME: fix divide by zero panic
	if m.B == 0 {
		return 0
	}
	return m.A / m.B
}

// Multipart 乘法
func (m *Math) Multipart() int {
	// TO DO: implement multipart
	return m.A * m.B
}

// SetA 设置 A 的值
func (m *Math) SetA(a int) {
	m.A = a
}

// GetA 获取 A 的值
func (m *Math) GetA() int {
	return m.A
}

// SetB 设置 B 的值
func (m *Math) SetB(b int) {
	m.B = b
}

// GetB 获取 B 的值
func (m *Math) GetB() int {
	return m.B
}

// ToMap 转换成 Map
func (m *Math) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(*m)
	if err != nil {
		return nil, err
	}

	jsonObj, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}

	return jsonObj.Map()
}

package math

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Logf("new math: %d", New(2, 1))
}

func TestDivide(t *testing.T) {
	m := New(1, 1)
	m.Divide()
	m.SetB(0)
	m.Divide()
}

func TestMultipart(t *testing.T) {
	m := New(1, 2)
	m.SetA(2)
	m.SetB(3)
	t.Logf("A: %d, B: %d", m.GetA(), m.GetB())
	t.Logf("2 * 3 = %d", m.Multipart())
}

func TestToMap(t *testing.T) {
	m := New(0, 2)
	map1, err := m.ToMap()
	if err != nil {
		panic(err)
	}
	t.Logf("map:%s", map1)
}

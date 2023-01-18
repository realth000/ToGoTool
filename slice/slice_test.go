package slice

import (
	"reflect"
	"testing"
)

func TestCleanDuplicate(t *testing.T) {
	// TODO: Simplify this.
	t.Run("int", func(t *testing.T) {
		data1 := []int{1, 2, 1, -2, 2, 3, 4, -1}
		data2 := []int{1, 2, -2, 3, 4, -1}
		data3 := CleanDuplicate(data1)
		if !reflect.DeepEqual(data2, data3) {
			t.Errorf("not equal, expected %v, got %v", data2, data3)
		}
	})
	t.Run("bool", func(t *testing.T) {
		data1 := []bool{true, true, false}
		data2 := []bool{true, false}
		data3 := CleanDuplicate(data1)
		if !reflect.DeepEqual(data2, data3) {
			t.Errorf("not equal, expected %v, got %v", data2, data3)
		}
	})
	t.Run("float32", func(t *testing.T) {
		data1 := []float32{1.00, 1.00, 2, 2, -1.10, -1.1, -2}
		data2 := []float32{1.00, 2, -1.1, -2}
		data3 := CleanDuplicate(data1)
		if !reflect.DeepEqual(data2, data3) {
			t.Errorf("not equal, expected %v, got %v", data2, data3)
		}
	})
	t.Run("string", func(t *testing.T) {
		data1 := []string{"a1", "a1", "1", "1", "", "", "\n", "\n", "a2"}
		data2 := []string{"a1", "1", "", "\n", "a2"}
		data3 := CleanDuplicate(data1)
		if !reflect.DeepEqual(data2, data3) {
			t.Errorf("not equal, expected %v, got %v", data2, data3)
		}
	})
	t.Run("pointer", func(t *testing.T) {
		p1, p2, p3, p4, p5 := 1, 2, 3, 4, 5
		data1 := []*int{&p1, &p1, &p2, &p3, &p4, &p2, &p4, &p5}
		data2 := []*int{&p1, &p2, &p3, &p4, &p5}
		data3 := CleanDuplicate(data1)
		if !reflect.DeepEqual(data2, data3) {
			t.Errorf("not equal, expected %v, got %v", data2, data3)
		}
	})
}

func TestByteFromString(t *testing.T) {
	for _, v := range []struct {
		name string
		data string
	}{
		{
			name: "normal",
			data: "running normal text",
		},
		{
			name: "multiline",
			data: "running \n multiline text",
		},
		{
			name: "complex",
			data: "rwe12QWE3中文  \n\t$-/*._\\",
		},
		{
			name: "string ``",
			data: `data
			asd 中 \\ \*&128`,
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			d1 := []byte(v.data)
			d2 := ByteFromString(v.data)
			if !reflect.DeepEqual(d1, d2) {
				t.Errorf("not equal, expected %v, got %v", d1, d2)
			}
		})
	}
}

func TestByteToString(t *testing.T) {
	for _, v := range []struct {
		name string
		data []byte
	}{
		{
			name: "normal",
			data: []byte("running normal text"),
		},
		{
			name: "multiline",
			data: []byte("running \n multiline text"),
		},
		{
			name: "complex",
			data: []byte("rwe12QWE3中文  \n\t$-/*._\\"),
		},
		{
			name: "string ``",
			data: []byte(`data
			asd 中 \\ \*&128`),
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			d1 := string(v.data)
			d2 := ByteToString(v.data)
			if !reflect.DeepEqual(d1, d2) {
				t.Errorf("not equal, expected %v, got %v", d1, d2)
			}
		})
	}
}

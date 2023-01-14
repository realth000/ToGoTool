package codec

import (
	"testing"
)

var (
	// UTF8 encoded "789zxc.*+测试优替哎福8"
	u    = []byte{55, 56, 57, 122, 120, 99, 46, 42, 43, 230, 181, 139, 232, 175, 149, 228, 188, 152, 230, 155, 191, 229, 147, 142, 231, 166, 143, 56}
	g    = []byte{55, 56, 57, 122, 120, 99, 46, 42, 43, 178, 226, 202, 212, 211, 197, 204, 230, 176, 165, 184, 163, 56}
	uBom = []byte{239, 187, 191, 55, 56, 57, 122, 120, 99, 46, 42, 43, 230, 181, 139, 232, 175, 149, 228, 188, 152, 230, 155, 191, 229, 147, 142, 231, 166, 143, 56}
)

func compareByteSlice(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	pos := len(a) - 1
	for pos >= 0 {
		if a[pos] != b[pos] {
			return false
		}
		pos--
	}
	return true
}

func TestIsUtf8(t *testing.T) {
	if IsUTF8(uBom) {
		t.Errorf("TestIsUtf8: expect false on UTF-8-BOM codec, but got true")
	}
	if IsUTF8(g) {
		t.Errorf("TestIsUtf8: expect false on GBK codec, but got true")
	}
	if !IsUTF8(u) {
		t.Errorf("TestIsUtf8: expect true on UTF-8 codec, but got false")
	}
}

func TestIsUtf8Bom(t *testing.T) {
	if IsUTF8BOM(u) {
		t.Errorf("TestIsUtf8BOM: expect false on UTF-8 codec, but got true")
	}
	if IsUTF8BOM(g) {
		t.Errorf("TestIsUtf8BOM: expect false on GBK codec, but got true")
	}
	if !IsUTF8BOM(uBom) {
		t.Errorf("TestIsUtf8BOM: expect true on UTF-8-BOM codec, but got false")
	}
}

func TestGbkToUtf8(t *testing.T) {
	s, err := GBKToUTF8(g)
	if err != nil {
		t.Errorf("TestGbkToUtf8: error transform: %v", err)
	}
	if !compareByteSlice(s, u) {
		t.Errorf("TestGbkToUtf8: transform result compare failed:\nresult is %v\nshould be %v", s, u)
	}
}

func TestUtf8ToGbk(t *testing.T) {
	s, err := UTF8ToGBK(u)
	if err != nil {
		t.Errorf("TestUtf8ToGbk: error transform: %v", err)
	}
	if !compareByteSlice(s, g) {
		t.Errorf("TestUtf8ToGbk: transform result compare failed:\nresult is %v\nshould be %v", s, g)
	}
	s, err = UTF8ToGBK(uBom)
	if err != nil {
		t.Errorf("TestUtf8ToGbk: error transform: %v", err)
	}
	if !compareByteSlice(s, g) {
		t.Errorf("TestUtf8ToGbk: transform result compare failed:\nresult is %v\nshould be %v", s, g)
	}
}

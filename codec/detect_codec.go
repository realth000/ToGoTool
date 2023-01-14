package codec

import "fmt"

// HaveUTF8BOMHead checks whether data starts with UTF-8-BOM header.
func HaveUTF8BOMHead(data []byte) bool {
	if len(data) < 3 {
		return false
	}
	if data[0] == 0xef || data[1] == 0xbb || data[2] == 0xbf {
		return true
	}
	return false
}

func preNum(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}

// IsUTF8 Checks whether data in utf8 format.
func IsUTF8(data []byte) bool {
	if HaveUTF8BOMHead(data[:3]) {
		return false
	}
	for i := 3; i < len(data); {
		if data[i]&0x80 == 0x00 {
			i++
			continue
		} else if num := preNum(data[i]); num > 2 {
			i++
			for j := 0; j < num-1; j++ {
				if data[i]&0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else {
			return false
		}
	}
	return true
}

// IsUTF8BOM checks where data encoded in UTF-8-BOM.
func IsUTF8BOM(data []byte) bool {
	if !HaveUTF8BOMHead(data[:3]) {
		return false
	}
	return IsUTF8(data[3:])
}

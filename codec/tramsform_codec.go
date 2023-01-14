package codec

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// GBKStringToUTF8 converts GBK encoded string gbkString to UTF-8 string.
func GBKStringToUTF8(gbkString string) (string, error) {
	utf8Bytes, err := GBKToUTF8([]byte(gbkString))
	return string(utf8Bytes), err
}

// GBKToUTF8 converts GBK encoded byte slice gbkBytes to UTF-8 byte slice.
func GBKToUTF8(gbkBytes []byte) ([]byte, error) {
	gbkReader := bytes.NewReader(gbkBytes)
	utfReader := transform.NewReader(gbkReader, simplifiedchinese.GBK.NewDecoder())
	utfBytes, err := io.ReadAll(utfReader)
	if err != nil {
		return nil, err
	}
	return utfBytes, nil
}

// UTF8StringToGBK converts UTF-8 encoded string utf8String to GBK string.
func UTF8StringToGBK(utf8String string) (string, error) {
	gbkBytes, err := UTF8ToGBK([]byte(utf8String))
	return string(gbkBytes), err
}

// UTF8ToGBK converts UTF-8 encoded byte slice utfBytes to GBK byte slice.
func UTF8ToGBK(utfBytes []byte) ([]byte, error) {
	var utfReader *bytes.Reader
	if HaveUTF8BOMHead(utfBytes) {
		utfReader = bytes.NewReader(utfBytes[3:])
	} else {
		utfReader = bytes.NewReader(utfBytes)
	}
	gbkReader := transform.NewReader(utfReader, simplifiedchinese.GBK.NewEncoder())
	gbkBytes, err := io.ReadAll(gbkReader)
	if err != nil {
		return nil, err
	}
	return gbkBytes, nil
}

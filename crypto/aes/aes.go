package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

/*
	ECB should not be used if encrypting more than one block of data with the same key.
	ModeCBC, OFB and CFB are similar, however OFB/CFB is better because you only need encryption and not decryption, which can save code space.
	CTR is used if you want good parallelization (ie. speed), instead of ModeCBC/OFB/CFB.
	XTS mode is the most common if you are encoding a random accessible data (like a hard disk or RAM).
	OCB is by far the best mode, as it allows encryption and authentication in a single pass. However there are patents on it in USA.
*/

type Mode = int8

const (
	ModeCBC Mode = iota
	ModeCFB
	ModeOFB
	ModeCTR
)

type Type = int8

const (
	TypeInvalid      = -1
	Type128     Type = iota
	Type192
	Type256
)

func checkAESKeyLength(key []byte) error {
	switch len(key) {
	case 16, 24, 32:
		return nil
	default:
		return fmt.Errorf("invalid AES key length:%d", len(key))
	}
}

func checkAESType(key []byte) Type {
	switch len(key) {
	case 16:
		return Type128
	case 24:
		return Type192
	case 32:
		return Type256
	default:
		return TypeInvalid
	}
}

func GenerateAESKey(size Type) ([]byte, error) {
	var s int
	switch size {
	case Type128:
		s = 16
	case Type192:
		s = 24
	case Type256:
		s = 32
	default:
		return nil, fmt.Errorf("invalid AES key size: %d", size)
	}
	ret := make([]byte, s)
	if _, err := rand.Read(ret); err != nil {
		return nil, fmt.Errorf("failed to generate: %v", err)
	}
	return ret, nil
}

func Encrypt(mode Mode, key []byte, data []byte) ([]byte, error) {
	aesType := checkAESType(key)
	if aesType == TypeInvalid {
		return nil, fmt.Errorf("invalid AES key length: %d", len(key))
	}
	data = addAESPadding(data)
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("too short cipher data with length %d", len(data))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherData := make([]byte, len(data)+aes.BlockSize)
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	switch mode {
	case ModeCBC:
		blockMode := cipher.NewCBCEncrypter(block, iv)
		blockMode.CryptBlocks(cipherData[aes.BlockSize:], data)
	case ModeCFB:
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(cipherData[aes.BlockSize:], data)
	case ModeOFB:
		stream := cipher.NewOFB(block, iv)
		stream.XORKeyStream(cipherData[aes.BlockSize:], data)
	case ModeCTR:
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(cipherData[aes.BlockSize:], data)
	default:
		return nil, fmt.Errorf("unknown AES mode: %d", mode)
	}
	return cipherData, nil
}

func Decrypt(mode Mode, key []byte, data []byte) ([]byte, error) {
	aesType := checkAESType(key)
	if aesType == TypeInvalid {
		return nil, fmt.Errorf("invalid AES key length: %d", len(key))
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("too short cipher data with length: %d", len(data))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failde to create cipher: %v", err)
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("invalid cipher data length after remove iv: %d", len(data))
	}
	//plainData := make([]byte, len(data)-aes.BlockSize)
	plainData := make([]byte, len(data))
	switch mode {
	case ModeCBC:
		blockMode := cipher.NewCBCDecrypter(block, iv)
		blockMode.CryptBlocks(plainData, data)
	case ModeCFB:
		blockMode := cipher.NewCFBDecrypter(block, iv)
		blockMode.XORKeyStream(plainData, data)
	case ModeOFB:
		stream := cipher.NewOFB(block, iv)
		stream.XORKeyStream(plainData, data)
	case ModeCTR:
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plainData, data)
	default:
		return nil, fmt.Errorf("unknown AES mode: %d", mode)
	}
	return removeAESPadding(plainData), nil
}

func addAESPadding(data []byte) []byte {
	paddingLength := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, padding...)
}

func removeAESPadding(data []byte) []byte {
	paddingLength := data[len(data)-1]
	return data[:len(data)-int(paddingLength)]
}

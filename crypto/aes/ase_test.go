package aes

import (
	"testing"
)

func TestGenerateAESKey(t *testing.T) {
	for _, v := range []struct {
		aesType      Type
		aesTypeName  string
		aesKeyLength int
	}{
		{
			aesType:      Type128,
			aesTypeName:  "AES-128",
			aesKeyLength: 16,
		},
		{
			aesType:      Type192,
			aesTypeName:  "AES-192",
			aesKeyLength: 24,
		},
		{
			aesType:      Type256,
			aesTypeName:  "AES-256",
			aesKeyLength: 32,
		},
	} {
		key1, err := GenerateAESKey(v.aesType)
		if err != nil {
			t.Errorf("error generating %s key: %v", v.aesTypeName, err)
		}
		if len(key1) != v.aesKeyLength {
			t.Errorf("invalid %s key: %v", v.aesTypeName, err)
		}
	}
}

func TestEncrypt(t *testing.T) {
	for _, v := range []struct {
		aesType     Type
		aesTypeName string
		data        []byte
	}{
		{
			aesType:     Type128,
			aesTypeName: "AES-128",
			data:        []byte("testTEST1478963?$#//.+&\\+/-_:测试"),
		},
		{
			aesType:     Type128,
			aesTypeName: "AES-128",
			data:        []byte("testTEST1478963?$#//.+&\\+/-_:测试"),
		},
		{
			aesType:     Type128,
			aesTypeName: "AES-128",
			data:        []byte("testTEST1478963?$#//.+&\\+/-_:测试"),
		},
	} {
		for _, m := range []struct {
			mode     Mode
			modeName string
		}{
			{
				mode:     ModeCBC,
				modeName: "CBC",
			},
			{
				mode:     ModeCFB,
				modeName: "CFB",
			},
			{
				mode:     ModeOFB,
				modeName: "OFB",
			},
			{
				mode:     ModeCTR,
				modeName: "CTR",
			},
		} {
			key, err := GenerateAESKey(v.aesType)
			if err != nil {
				t.Errorf("failed to generate %s key in mode %s: %v", v.aesTypeName, m.modeName, err)
			}
			cipherData, err := Encrypt(m.mode, key, v.data)
			if err != nil {
				t.Errorf("failed to encrty %s in mode %s: %v", v.aesTypeName, m.modeName, err)
			}

			plainData, err := Decrypt(m.mode, key, cipherData)
			if err != nil {
				t.Errorf("failed to decrypt %s in mode %s: %v", v.aesTypeName, m.modeName, err)
			}
			if len(v.data) != len(plainData) {
				t.Errorf("failed check for %s in mode %s: not equal data length after decrypt, expected %d, got %d",
					v.aesTypeName, m.modeName, len(v.data), len(plainData))
			}
			for i, c := range v.data {
				if c != plainData[i] {
					t.Errorf("failed check for %s in mode %s: not equal data at index %d, expected %d, got %d",
						v.aesTypeName, m.modeName, i, c, plainData[i])
				}
			}
		}
	}

}

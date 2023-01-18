package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
)

type SumType = int8

const (
	SumMD5 SumType = iota
	SumSHA1
	SumSHA224
	SumSHA256
	SumSHA384
	SumSHA512
	SumSHA512_256
	SumSHA3_224
	SumSHA3_256
	SumSHA3_384
	SumSHA3_512
)

func MD5(data []byte) [16]byte {
	return md5.Sum(data)
}

func MD5String(data string) [16]byte {
	return MD5([]byte(data))
}

func SHA1(data []byte) [20]byte {
	return sha1.Sum(data)
}

func SHA1String(data string) [20]byte {
	return SHA1([]byte(data))
}

func SHA224(data []byte) [28]byte {
	return sha256.Sum224(data)
}

func SHA224String(data string) [28]byte {
	return SHA224([]byte(data))
}

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func SHA256String(data string) [32]byte {
	return SHA256([]byte(data))
}

func SHA384(data []byte) [48]byte {
	return sha512.Sum384(data)
}

func SHA384String(data string) [48]byte {
	return SHA384([]byte(data))
}

func SHA512(data []byte) [64]byte {
	return sha512.Sum512(data)
}

func SHA512String(data string) [64]byte {
	return SHA512([]byte(data))
}

func SHA512_224(data []byte) [28]byte {
	return sha3.Sum224(data)
}

func SHA512_224String(data string) [28]byte {
	return SHA512_224([]byte(data))
}

func SHA512_256(data []byte) [32]byte {
	return sha512.Sum512_256(data)
}

func SHA512_256String(data string) [32]byte {
	return SHA512_256([]byte(data))
}

func SHA3_224(data []byte) [28]byte {
	return sha3.Sum224(data)
}

func SHA3_224String(data string) [28]byte {
	return SHA3_224([]byte(data))
}

func SHA3_256(data []byte) [32]byte {
	return sha3.Sum256(data)
}

func SHA3_256String(data string) [32]byte {
	return SHA3_256([]byte(data))
}

func SHA3_384(data []byte) [48]byte {
	return sha3.Sum384(data)
}

func SHA3_384String(data string) [48]byte {
	return SHA3_384([]byte(data))
}

func SHA3_512(data []byte) [64]byte {
	return sha3.Sum512(data)
}

func SHA3_512String(data string) [64]byte {
	return sha3.Sum512([]byte(data))
}

func Hash(sumType SumType, data []byte) []byte {
	switch sumType {
	case SumMD5:
		v := MD5(data)
		return v[:]
	case SumSHA1:
		v := SHA1(data)
		return v[:]
	case SumSHA224:
		v := SHA224(data)
		return v[:]
	case SumSHA256:
		v := SHA256(data)
		return v[:]
	case SumSHA384:
		v := SHA384(data)
		return v[:]
	case SumSHA512:
		v := SHA512(data)
		return v[:]
	case SumSHA512_256:
		v := SHA512_256(data)
		return v[:]
	case SumSHA3_224:
		v := SHA3_224(data)
		return v[:]
	case SumSHA3_256:
		v := SHA3_256(data)
		return v[:]
	case SumSHA3_384:
		v := SHA3_384(data)
		return v[:]
	case SumSHA3_512:
		v := SHA3_512(data)
		return v[:]
	default:
		panic("invalid sum hash type")
	}
}

func HashString(sumType SumType, data string) []byte {
	switch sumType {
	case SumMD5:
		v := MD5String(data)
		return v[:]
	case SumSHA1:
		v := SHA1String(data)
		return v[:]
	case SumSHA224:
		v := SHA224String(data)
		return v[:]
	case SumSHA256:
		v := SHA256String(data)
		return v[:]
	case SumSHA384:
		v := SHA384String(data)
		return v[:]
	case SumSHA512:
		v := SHA512String(data)
		return v[:]
	case SumSHA512_256:
		v := SHA512_256String(data)
		return v[:]
	case SumSHA3_224:
		v := SHA3_224String(data)
		return v[:]
	case SumSHA3_256:
		v := SHA3_256String(data)
		return v[:]
	case SumSHA3_384:
		v := SHA3_384String(data)
		return v[:]
	case SumSHA3_512:
		v := SHA3_512String(data)
		return v[:]
	default:
		panic("invalid sum hash type")
	}
}

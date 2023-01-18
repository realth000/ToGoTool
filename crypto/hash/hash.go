package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func SHA256String(data string) [32]byte {
	return SHA256([]byte(data))
}

func SHA512256(data []byte) [32]byte {
	return sha512.Sum512_256(data)
}

func SHA512256String(data string) [32]byte {
	return SHA512256([]byte(data))
}

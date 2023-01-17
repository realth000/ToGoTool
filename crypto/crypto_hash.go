package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func HashMd5(data []byte) [16]byte {
	return md5.Sum(data)
}

func HashMd5String(data string) [16]byte {
	return HashMd5([]byte(data))
}

func HashSha1(data []byte) [20]byte {
	return sha1.Sum(data)
}

func HashSha1String(data string) [20]byte {
	return HashSha1([]byte(data))
}

func HashSha256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func HashSha256String(data string) [32]byte {
	return HashSha256([]byte(data))
}

func HashSha512256(data []byte) [32]byte {
	return sha512.Sum512_256(data)
}

func HashSha512256String(data string) [32]byte {
	return HashSha512256([]byte(data))
}

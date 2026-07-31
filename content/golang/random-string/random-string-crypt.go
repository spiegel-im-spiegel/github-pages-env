package randstr

import (
	"crypto/rand"
	"unsafe"
)

type CryptoRandom struct{}

func NewCryptoRandom() Random {
	return &CryptoRandom{}
}

func (cr *CryptoRandom) String(len int) string {
	b := make([]byte, len)
	for i, offset, size, rest := 0, 0, 0, 0; i < len; rest-- {
		//fmt.Println(i, offset, size, rest)
		if rest <= 0 {
			offset = i
			var err error
			size, err = rand.Read(b[offset:])
			if err != nil || size <= 0 {
				return ""
			}
			rest = size
		}
		if idx := int(b[offset+(size-rest)] & letterIdxMask); idx < letterBytesLen {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Refer:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// https://zenn.dev/najeira/articles/2017-12-28-qiita-5974c6b2c59ecc02fad2
package randstr

import (
	"math/rand"
	"unsafe"
)

type MathRandom struct {
	src rand.Source
}

func NewMathRandom(seed int64) Random {
	return &MathRandom{src: rand.NewSource(seed)}
}

func (mr *MathRandom) String(len int) string {
	b := make([]byte, len)
	for i, cache, rest := 0, mr.src.Int63(), letterIdxMax; i < len; rest-- {
		if rest <= 0 {
			cache, rest = mr.src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < letterBytesLen {
			b[i] = letterBytes[idx]
			i++
		}
		cache >>= letterIdxBits
	}
	return *(*string)(unsafe.Pointer(&b))
}

// func (mr *MathRandom) String2(len int) string {
// 	b := make([]byte, len)
// 	for i, offset, size, rest := 0, 0, 0, 0; i < len; rest-- {
// 		//fmt.Println(i, offset, size, rest)
// 		if rest <= 0 {
// 			offset = i
// 			var err error
// 			size, err = mr.Read(b[offset:])
// 			if err != nil || size <= 0 {
// 				return ""
// 			}
// 			rest = size
// 		}
// 		if idx := int(b[offset+(size-rest)] & letterIdxMask); idx < letterBytesLen {
// 			b[i] = letterBytes[idx]
// 			i++
// 		}
// 	}
// 	return *(*string)(unsafe.Pointer(&b))
// }

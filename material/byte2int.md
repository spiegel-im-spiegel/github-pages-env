# バイト列を整数に変換する簡単なお仕事メモ

簡単な記述を教えてもらったので，忘れないうちにメモ。

バイト列を整数に変換する簡単なお仕事。

```go
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	octets := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}

	fmt.Printf("%#016x\n", binary.BigEndian.Uint64(octets))    //0x0001020304050607
	fmt.Printf("%#016x\n", binary.LittleEndian.Uint64(octets)) //0x0706050403020100
}
```

整数からバイト列への変換も簡単。

```go
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	value := uint64(0x0706050403020100)
	buf := make([]byte, binary.MaxVarintLen64) //MaxVarintLen64 = 10

	binary.BigEndian.PutUint64(buf, value)
	fmt.Println(buf[:8]) //[7 6 5 4 3 2 1 0]
	binary.LittleEndian.PutUint64(buf, value)
	fmt.Println(buf[:8]) //[0 1 2 3 4 5 6 7]

}
```

なんちうか今まですンごい回りくどい記述してたよ。わざわざ `bytes.NewReader(octets)` でラップするとか。

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	octets := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	var value uint64

	if err := binary.Read(bytes.NewReader(octets), binary.LittleEndian, &value); err != nil {
		return
	}
	fmt.Printf("%#016x\n", value) //0x0706050403020100
}
```

浮動小数点数に変換するならともかく，整数なんだから... orz
明日からはちゃんとやろう。

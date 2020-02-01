package calculator

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// + 号拼接
	// fmt.Sprintf
	// strings.Join
	// buffer.WriteString
/*
	 在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
	 在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
	 较少字符串连接的场景下性能最好，而且代码更简短清晰，可读性更好
	 如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf
*/
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}

	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}

	var buffer bytes.Buffer
	for i := 0; i < 100; i++ {
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
	}
	_ = buffer.String()
}

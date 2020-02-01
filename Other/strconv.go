package calculator

import (
	"fmt"
	"strconv"
)

func toString(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var n = 1000
	s := toString(n)
	fmt.Println(s)
	s1 := Reverse(s)
	s2,_ := strconv.ParseInt(s1, 10, 0)
	fmt.Println(s2 == 1)

	fmt.Println(strconv.Atoi("1000"))
}

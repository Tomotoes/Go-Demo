package calculator

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	return regexp.MustCompile("[-_](.)").ReplaceAllStringFunc(s, func(w string) string {
		fmt.Println(w)
		return strings.ToUpper(w[1:])
	})
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
}

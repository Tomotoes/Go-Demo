package calculator

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	FILENAME := "C:/Users/jinma/Desktop/毒鸡汤.txt"
	file, err := os.Open(FILENAME)
	check(err)
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		word, err := br.ReadString(byte('\n'))
		if err != nil {
			return
		}
		fmt.Println(word)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

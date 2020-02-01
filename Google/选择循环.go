package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func grade(score float32) string {
	result := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		result = "D"
	case score < 70:
		result = "C"
	case score < 90:
		result = "B"
	case score <= 100:
		result = "A"
	}
	return result
}

func simpleEval(str string) float64 {
	resolve := strings.Split(strings.Trim(regexp.MustCompile(`\s+`).ReplaceAllString(str, " "), " "), " ")
	fmt.Println(resolve)
	a, _ := strconv.ParseFloat(resolve[0], 10)
	op := resolve[1]
	b, _ := strconv.ParseFloat(resolve[2], 10)
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "**":
		result = math.Pow(a, b)
	default:
		panic(fmt.Sprintf("Unsupported operator: %s\n", op))
	}
	return result
}

func convertBin(n int) string {
	result := ""
	for ; n > 2; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever")
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	const filename = "D:/Go/src/Google/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		log.Fatal(err)
	}
	fmt.Println(
		simpleEval("1 + 2"),
		simpleEval("2 **    5"),
		simpleEval("1 - -5"),
		simpleEval("1 * -8"),
		simpleEval("1 / 5"),
	)

	fmt.Println(
		grade(12),
		grade(56),
		grade(0),
		grade(100),
	)

	fmt.Println(
		convertBin(5),
		convertBin(13),
		convertBin(100),
	)

	printFile(filename)

	//forever()

	fmt.Println(apply(pow, 3, 4))

}

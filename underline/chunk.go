package underline

import _ "underline"

func Chunk(arr []interface{}, n int) [][]interface{} {
	length := len(arr)
	size := length/n + _.ToInt(length%n != 0)
	chunks := make([][]interface{}, size)
	for i := 0; i < size; i++ {
		chunks[i] = arr[i*size : i*size+size]
	}
	return chunks
}

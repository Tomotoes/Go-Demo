package main

func Index(arr []string, target string) int {
	for idx, value := range arr {
		if value == target {
			return idx
		}
	}
	return -1
}

func Include(arr []string, target string) bool {
	return Index(arr, target) != -1
}

func Any(arr []string, f func(string) bool) bool {
	for _, s := range arr {
		if f(s) {
			return true
		}
	}
	return false
}

func All(arr []string, f func(string) bool) bool {
	for _, s := range arr {
		if !f(s) {
			return false
		}
	}
	return true
}

func Filter(arr []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, s := range arr {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func Map(arr []string, f func(string) string) []string {
	result := make([]string, 0)
	for _, s := range arr {
		result = append(result, f(s))
	}
	return result
}

func main() {
	// TODO: 测试~
}

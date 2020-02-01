package underline

func ToInt(n interface{}) int {
	switch n.(type) {
	case bool:
		t := n.(bool)
		if t {
			return 1
		} else {
			return 0
		}
	}
}

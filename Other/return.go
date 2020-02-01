package main

func getFirstName() (string, error) {
	return "simon", nil
}

func Name() (name string) {
	name = "leann"
	if name, err = getFirstName(); err == nil {
		return
	}
	return
}

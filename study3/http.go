package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")

	t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.data}}</div>")

	t.Execute(w, map[string]string{"name": "Simon", "data": "helloWorld"})
}

func main() {
	http.HandleFunc("/", handleHTTP)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}

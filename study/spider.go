package main

type Fetcher interface {
	Fetch(url string)(body string , urls []string ,err error)
}

func main() {

}
package main

import "fmt"

func getNotification(user string) chan string {
	notification := make(chan string)
	go func() {
		notification <- fmt.Sprintf("Hi %s,welcome to tomotoes.com~", user)
	}()

	return notification
}

func main() {
	jack, simon := getNotification("jack"), getNotification("simon")
	fmt.Println(<-jack)
	fmt.Println(<-simon)
}

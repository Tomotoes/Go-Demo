package main

import "fmt"

type IAuthor interface {
	fullName() string
}

type author struct {
	firstName string
	lastName  string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type IPost interface {
	details()
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
}

type website struct {
	posts []post
}

func (w website) contents() {
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{firstName: "Simon", lastName: "Ma"}
	post1 := post{"Go", "Go is less!", author1}
	post2 := post{"Ao", "Ao is less!", author1}
	post3 := post{"Bo", "Bo is less!", author1}

	w := website{posts: []post{post1, post2, post3}}

	w.contents()
}

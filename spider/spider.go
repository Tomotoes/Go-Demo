package main

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	URL      = "http://www.nows.fun/"
	SELECTOR = "#sentence"
	REQUESTS = 99
	FILENAME = "C:/Users/jinma/Desktop/毒鸡汤.txt"
)

var WG = sync.WaitGroup{}
var Lock = sync.RWMutex{}
var file *os.File
var Words = map[string]bool{}
var collect = ""

func init() {
	var err error
	defer check(err)
	if checkFileIsExist(FILENAME) {
		file, err = os.OpenFile(FILENAME, os.O_APPEND, 0666)
		br := bufio.NewReader(file)
		for {
			word, err := br.ReadString(byte('\n'))
			if err != nil {
				return
			}
			Words[word] = true
		}
	} else {
		file, err = os.Create(FILENAME)
	}
}

func main() {
	for i := 0; i < REQUESTS; i++ {
		go func() {
			time.Sleep(time.Millisecond * 500)

			WG.Add(1)
			defer WG.Done()
			var word = Spider()
			Lock.RLock()
			existing := Words[word]
			Lock.RUnlock()
			if !existing {
				Lock.Lock()
				collect += word + "\n"
				Words[word] = true
				Lock.Unlock()
			}
		}()
	}
	WG.Wait()
	file.WriteString(collect)
	file.Close()
}

func Spider() string {
	res, err := http.Get(URL)
	check(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		check(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	check(err)
	return doc.Find(SELECTOR).Text()
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

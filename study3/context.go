package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	var (
		urls        = []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
		wg          = sync.WaitGroup{}
		ctx, cancel = context.WithCancel(context.Background())
	)
	for _, url := range urls {
		wg.Add(1)
		subCtx := context.WithValue(ctx, "url", url)
		go request(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func request(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value("url").(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop getting url: ", url)
			return
		default:
			resp, err := http.Get(url)
			if err == nil && resp.StatusCode == http.StatusOK {
				body, _ := ioutil.ReadAll(resp.Body)
				subCtx := context.WithValue(ctx, "resp", fmt.Sprintf("%s:%x", url, md5.Sum(body)))
				wg.Add(1)
				go show(subCtx, wg)
			}
			_ = resp.Body.Close()
			time.Sleep(time.Second * 1)
		}
	}
}

func show(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp")
			return
		default:
			// do something , ex: orm db or rpc
			fmt.Println("printing: ", ctx.Value("resp").(string))
			time.Sleep(time.Second * 1)
		}
	}
}

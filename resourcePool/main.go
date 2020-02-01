package main

import (
	"fmt"
	"log"
	"math/rand"
	"other/resourcePool/db"
	"other/resourcePool/pool"
	"sync"
	"time"
)

const (
	maxGoruntine   = 5
	pooledResource = 2
)

func query(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		fmt.Println("Acquire conn error, ", err)
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	log.Printf("QID[%d] CID[%d]\n", id, conn.(*db.DBConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoruntine)

	p, err := pool.New(db.CreateConn, pooledResource)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < maxGoruntine; i++ {
		go func(id int) {
			query(i, p)
			wg.Done()
		}(i)
	}
	wg.Wait()
	p.Close()
	fmt.Println("finish")
}

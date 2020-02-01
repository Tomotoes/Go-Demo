package db

import (
	"fmt"
	"io"
	"sync/atomic"
)

var idCounter int32

type DBConnection struct {
	ID int32
}

func (conn *DBConnection) Close() error {
	fmt.Println("conn closed")
	return nil
}

func CreateConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("create conn ,id: ", id)
	return &DBConnection{
		ID: id,
	}, nil
}

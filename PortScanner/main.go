package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// command param
var (
	host            string
	start           int
	end             int
	timeout         string
	timeoutDuration time.Duration
	pause           string
	pauseDuration   time.Duration
	log             bool
)

// port util
var (
	minPort = 80
	maxPort = 65535
	isPort  = func(port int) bool {
		return port >= minPort && port <= maxPort
	}
)

// custom error
var (
	startPortError                     = fmt.Errorf("starting port out of range (should be between %d and %d)", minPort, maxPort)
	endPortError                       = fmt.Errorf("ending port out of range (should be between %d and %d)", minPort, maxPort)
	endPortIsSmallerThanStartPortError = errors.New("ending port must be greater than starting port")
)

// custom type
type status string

func (s *status) init(isOpened bool) {
	if isOpened {
		*s = "OPENED"
	} else {
		*s = "CLOSED"
	}
}

func (s status) isOpened() bool {
	return s == "OPENED"
}

func (s status) isClosed() bool {
	return s == "CLOSED"
}

// global variable
var (
	openedPorts []int
	wg          sync.WaitGroup
	lock        sync.Mutex
)

func init() {
	flag.StringVar(&host, "host", "localhost", "the host to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "timeout (e.g. 50ms or 1s)")
	flag.StringVar(&pause, "pause", "1ms", "pause after each scanned port (e.g. 5ms)")
	flag.IntVar(&start, "start", minPort, "the lower end to scan")
	flag.IntVar(&end, "end", maxPort, "the upper end to scan")
	flag.BoolVar(&log, "log", true, "print result when scanning the port (true/false)")
	flag.Parse()

	if !isPort(start) {
		panic(startPortError)
	}
	if !isPort(end) {
		panic(endPortError)
	}
	if end < start {
		panic(endPortIsSmallerThanStartPortError)
	}
	var err error
	timeoutDuration, err = time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}

	pauseDuration, err = time.ParseDuration(pause)
	if err != nil {
		panic(err)
	}
}

func scan(port int) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeoutDuration)

	var s status

	isOpened := true
	if err, ok := err.(net.Error); ok && err.Timeout() {
		isOpened = false
	}
	s.init(isOpened)

	if log {
		fmt.Printf("scan: %d , status: %v\n", port, s)
	}
	if s.isOpened() {
		if conn != nil {
			conn.Close()
		}

		lock.Lock()
		openedPorts = append(openedPorts, port)
		lock.Unlock()
	}

}

func main() {
	startTime := time.Now()
	for port := start; port <= end; port++ {
		wg.Add(1)
		go scan(port)
		time.Sleep(pauseDuration)
	}
	wg.Wait()

	fmt.Printf("\nscan finished in %v\n", time.Since(startTime))

	if len(openedPorts) == 0 {
		fmt.Println("Ports are off.")
		return
	}

	fmt.Println("Openned ports: ")
	for _, port := range openedPorts {
		fmt.Println(port)
	}
}

package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var openPorts int = 0
var closedPorts int = 0

const (
	version = "1.0.0"
	author  = "Aditya Pawar"
)

// Port scanning function for scanning out the port
func scanPort(host string, port int, wg *sync.WaitGroup) {

	defer wg.Done() // For conforming that the task has been done.

	address := fmt.Sprintf("%s:%d", host, port)

	connection, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		//fmt.Printf("Port %d is closed\n", port)
		closedPorts++
		return
	}

	connection.Close()

	openPorts++
	fmt.Printf("Port %d is open\n", port)
}

func main() {

	host := flag.String("host", "127.0.0.1", "Target host")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage : %s\n", "go run main.go -host \"Port\" -start startPort -end endPort")
		fmt.Fprintf(flag.CommandLine.Output(), "Version : %s\n", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Author : %s\n", author)
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	var wg sync.WaitGroup

	fmt.Printf("Scanning ports on : %s ...\n", *host)

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go scanPort(*host, port, &wg)
	}

	wg.Wait()

	fmt.Printf("Opened Ports are : %d \n", openPorts)
	fmt.Printf("Closed Ports are : %d \n", closedPorts)

	fmt.Println("Scan completed.")
}

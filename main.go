package main

import (
	"fmt"
	"github.com/davisfreimanis/go-scan/port"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(port.ScanPort(8000, "localhost", "tcp"))
	res := port.ScanWellKnownPorts("localhost")
	port.PrintScanResult(res)
}

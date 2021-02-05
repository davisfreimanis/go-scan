package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// ScanResult stores the values of a result
type ScanResult struct {
	port   int
	status string
}

// ScanPort scans a port
func ScanPort(port int, hostname string, protocol string) ScanResult {
	result := ScanResult{port: port}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Second)

	if err != nil {
		result.status = "Closed"
		return result
	}

	defer conn.Close()

	result.status = "Open"
	return result
}

// ScanWellKnownPorts scans port 1-1024
func ScanWellKnownPorts(hostname string) []ScanResult {
	results := []ScanResult{}

	for i := 1; i <= 1024; i++ {
		res := ScanPort(i, hostname, "tcp")
		results = append(results, res)
	}

	return results
}

// PrintScanResult prints the result from a scan
func PrintScanResult(scanResult []ScanResult) {
	fmt.Println("Printing open ports")
	for _, r := range scanResult {
		if r.status == "Open" {
			fmt.Println(r.port, " is ", r.status)
		}
	}
}

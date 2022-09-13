package details

import (
	"log"
	"net"
	"os"
)

// uses the os package to show the hostname
func GetHostName() (string, error) {
	hostname, _ := os.Hostname()
	return hostname, nil
}

// uses the net package to show the IP address
func GetIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}

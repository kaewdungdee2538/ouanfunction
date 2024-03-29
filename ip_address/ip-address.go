package ip_address

import (
	"fmt"
	"net"
	"os"
)

func GetLocalIpAddress() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			return ipv4.String()
		}
	}
	return ""
}

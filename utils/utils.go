// Package utils .
// Contains utilities that is used throughout the application
package utils

import (
	"net"

	"github.com/VinukaThejana/go-utils/logger"
)

var log logger.Logger

// GetOutboundIP is a function that is used to get the IP assigned to the interface
// that can connect to the internet
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Errorf(err, nil)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

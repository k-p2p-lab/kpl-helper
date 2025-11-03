package address

import "fmt"

// IPv4 represents an IPv4 address with its port.
type IPv4 struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// String returns the string representation of the IPv4 address in "address:port" format.
func (ip *IPv4) String() string {
	return fmt.Sprintf("%s:%d", ip.Address, ip.Port)
}

// NewIPv4 creates a new IPv4 address instance.
func NewIPv4(address string, port int) *IPv4 {
	return &IPv4{
		Address: address,
		Port:    port,
	}
}

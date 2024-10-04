package main
import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	parts := make([]string, 4)
	for i, b := range ip {
		parts[i] = fmt.Sprintf("%d", b)
	}
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 0},
		"Google": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

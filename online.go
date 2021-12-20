package ami

import "net"

func Online() bool {
	if _, err := net.Dial("tcp", "8.8.8.8:53"); err == nil {
		return true
	}
	return false
}

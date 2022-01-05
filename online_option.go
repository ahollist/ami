package ami

import "time"

type dialConfig struct {
	network string
	address string
	port    string
	timeout time.Duration
}

var defaultConfig = dialConfig{
	network: "tcp",
	address: "8.8.8.8",
	port:    "53",
	timeout: 0,
}

func (dc dialConfig) AddressAndPortString() string {
	return dc.address + ":" + dc.port
}

func (dc dialConfig) Network() string {
	return dc.network
}

func (dc dialConfig) Timeout() time.Duration {
	return dc.timeout
}

type onlineOption func(*dialConfig)

func WithNetwork(network string) onlineOption {
	switch network {
	// IP and IPv6 is unsupported at the moment
	case "tcp", "tcp4", "udp", "udp4":
		break
	default:
		// todo: use an unsupported option type to inform of bad configurations
		return nil
	}
	return func(dc *dialConfig) {
		dc.network = network
	}
}

func WithAddress(address string) onlineOption {
	return func(dc *dialConfig) {
		dc.address = address
	}
}

func WithPort(port string) onlineOption {
	return func(dc *dialConfig) {
		dc.port = port
	}
}

func WithTimeout(timeout time.Duration) onlineOption {
	return func(dc *dialConfig) {
		dc.timeout = timeout
	}
}

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

func (dc dialConfig) ExportDialParams() (network, addressAndPort string, timeout time.Duration) { // needs test
	// Resolve Address+Port String first
	// Update Network if required
	return dc.Network(), dc.AddressAndPortString(), dc.Timeout()
}

func (dc dialConfig) AddressAndPortString() string { // needs test
	// todo: parse address and set network/port information appropriately (e.g. "www.google.com:http")
	return dc.address + ":" + dc.port
}

func (dc dialConfig) Network() string {
	return dc.network
}

func (dc dialConfig) Timeout() time.Duration {
	return dc.timeout
}

type onlineOption func(*dialConfig)

func WithNetwork(network string) onlineOption { // needs test
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

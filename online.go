package ami

import "net"

type dialConfig struct {
	network string
	address string
	port    string
}

var defaultConfig = dialConfig{
	network: "tcp",
	address: "8.8.8.8",
	port:    "53",
}

type configOption func(*dialConfig)

func WithNetwork(network string) configOption {
	switch network {
	// IP and IPv6 is unsupported at the moment
	case "tcp", "tcp4", "udp", "udp4":
		break
	default:
		return nil
	}
	return func(dc *dialConfig) {
		dc.network = network
	}
}

func WithAddress(address string) configOption {
	return func(dc *dialConfig) {
		dc.address = address
	}
}

func WithPort(port string) configOption {
	return func(dc *dialConfig) {
		dc.port = port
	}
}

func Online(opts ...configOption) (bool, error) {
	conf := defaultConfig
	for _, opt := range opts {
		opt(&conf)
	}
	_, err := net.Dial(conf.network, conf.address+":"+conf.port)
	if err == nil {
		return true, nil
	}
	return false, err
}

package ami

import "net"

// Online takes a variadic number of options
// Online returns a boolean for the connection status and any applicable errors from the connection attempt
func Online(opts ...onlineOption) (bool, error) {
	conf := defaultConfig
	for _, opt := range opts {
		opt(&conf)
	}
	_, err := net.Dial(conf.Network(), conf.AddressAndPortString())
	if err == nil {
		return true, nil
	}
	return false, err
}

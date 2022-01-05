package ami

import "net"

// Online takes a variadic number of options
// Online returns a boolean for the connection status and any applicable errors from the connection attempt
func Online(opts ...onlineOption) (bool, error) {
	conf := defaultConfig
	for _, opt := range opts {
		if opt != nil {
			opt(&conf)
		}
	}
	_, err := net.DialTimeout(conf.Network(), conf.AddressAndPortString(), conf.Timeout())
	if err == nil {
		return true, nil
	}
	return false, err
}

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
	c, err := net.DialTimeout(conf.ExportDialParams()) // do I need to close this instead
	if err == nil {
		return true, nil
	}
	c.Close()
	return false, err
}

// needs test for:
// localhost
// 8.8.8.8:53
// Invalid info
//

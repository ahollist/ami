package ami

import "net"

// Online takes a variadic number of options
// Online returns a boolean for the connection status, after calling Online you can check for any errors with
func Online(opts ...onlineOption) bool {
	conf := defaultConfig
	for _, opt := range opts {
		if opt != nil {
			opt(&conf)
		}
	}
	c, err := net.DialTimeout(conf.ExportDialParams()) // do I need to close this instead
	if err == nil {
		return true
	}
	c.Close()
	return false
}

// needs test for:
// localhost
// 8.8.8.8:53
// Invalid info
//

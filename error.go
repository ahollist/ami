package ami

var err error

// Error returns the last error from any exported function in the ami package
func Error() error {
	return err
}

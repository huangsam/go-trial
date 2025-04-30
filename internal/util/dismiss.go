package util

// Dismiss calls a function and ignores its error.
//
// This is useful for functions that close a resource and return an error that
// we don't care about, such as closing a file or a network connection.
func Dismiss(f func() error) {
	_ = f()
}

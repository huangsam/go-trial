package util

func Dismiss(f func() error) {
	_ = f()
}

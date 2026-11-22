package performance

import "sync"

const maxPooledMessageCapacity = 1024

var messageBuffers = sync.Pool{
	New: func() any {
		buffer := make([]byte, 0, 64)
		return &buffer
	},
}

// FormatMessageWithPool formats a greeting using a reusable byte buffer.
// sync.Pool may discard its contents at any time, so it is only an optimization.
func FormatMessageWithPool(name string) string {
	buffer := messageBuffers.Get().(*[]byte)
	*buffer = append((*buffer)[:0], "Hello, "...)
	*buffer = append(*buffer, name...)
	*buffer = append(*buffer, '!')
	message := string(*buffer)

	if cap(*buffer) <= maxPooledMessageCapacity {
		messageBuffers.Put(buffer)
	}
	return message
}

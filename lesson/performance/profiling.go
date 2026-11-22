package performance

import "fmt"

// ProfileWorkload formats records and computes a checksum for profiling examples.
func ProfileWorkload(iterations int) uint64 {
	var checksum uint64
	for i := range iterations {
		record := fmt.Sprintf("record-%d", i)
		for j := 0; j < len(record); j++ {
			checksum = checksum*31 + uint64(record[j])
		}
	}
	return checksum
}

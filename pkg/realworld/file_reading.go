package realworld

import (
	"bufio"
	"os"

	"github.com/huangsam/go-trial/internal/cmd"
)

// ReadLines reads a file and returns its content as a slice of strings.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer cmd.Dismiss(file.Close)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

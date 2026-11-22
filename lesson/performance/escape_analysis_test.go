package performance_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/performance"
	"github.com/stretchr/testify/assert"
)

func TestEscapeAnalysisExamples(t *testing.T) {
	assert.Equal(t, 42, performance.LocalRecordValue(42))
	assert.Equal(t, 42, performance.NewEscapeRecord(42).Value)
}

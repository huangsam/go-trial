package practical_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/practical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileRoundTrip(t *testing.T) {
	profile, wireData, err := practical.ProfileRoundTrip()
	require.NoError(t, err)
	require.NotEmpty(t, wireData)
	require.NotNil(t, profile.Score)

	assert.Equal(t, int64(42), profile.Id)
	assert.Equal(t, "Ada", profile.Name)
	assert.True(t, profile.Active)
	assert.Equal(t, int32(0), *profile.Score)
	assert.Equal(t, []string{"engineer", "teacher"}, profile.Tags)
	assert.Equal(t, "ada@example.com", profile.GetEmail())
}

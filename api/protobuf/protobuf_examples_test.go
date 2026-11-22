package protobuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestProfileProtoRoundTrip(t *testing.T) {
	score := int32(0)
	profile := &Profile{
		Id:     42,
		Name:   "Ada",
		Active: true,
		Status: Profile_STATUS_ACTIVE,
		Address: &Profile_Address{
			Street: "1 Main St",
			City:   "London",
		},
		Tags:     []string{"engineer", "teacher"},
		Metadata: map[string]string{"team": "platform"},
		Contact:  &Profile_Email{Email: "ada@example.com"},
		Score:    &score,
	}

	data, err := proto.Marshal(profile)
	require.NoError(t, err)

	decoded := new(Profile)
	require.NoError(t, proto.Unmarshal(data, decoded))
	assert.True(t, proto.Equal(profile, decoded))
	require.NotNil(t, decoded.Score)
	assert.Zero(t, *decoded.Score)
}

func TestProfileOptionalScoreAbsent(t *testing.T) {
	data, err := proto.Marshal(&Profile{})
	require.NoError(t, err)

	decoded := new(Profile)
	require.NoError(t, proto.Unmarshal(data, decoded))
	assert.Nil(t, decoded.Score)
}

package practical

import (
	pb "github.com/huangsam/go-trial/api/protobuf"
	"google.golang.org/protobuf/proto"
)

// ProfileRoundTrip marshals and unmarshals an example Profile.
func ProfileRoundTrip() (*pb.Profile, []byte, error) {
	score := int32(0)
	profile := &pb.Profile{
		Id:     42,
		Name:   "Ada",
		Active: true,
		Status: pb.Profile_STATUS_ACTIVE,
		Address: &pb.Profile_Address{
			Street: "1 Main St",
			City:   "London",
		},
		Tags:     []string{"engineer", "teacher"},
		Metadata: map[string]string{"team": "platform"},
		Contact:  &pb.Profile_Email{Email: "ada@example.com"},
		Score:    &score,
	}

	wireData, err := proto.Marshal(profile)
	if err != nil {
		return nil, nil, err
	}

	decoded := new(pb.Profile)
	if err := proto.Unmarshal(wireData, decoded); err != nil {
		return nil, nil, err
	}
	return decoded, wireData, nil
}

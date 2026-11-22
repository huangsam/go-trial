package sub

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/huangsam/go-trial/lesson/practical"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

// ProtobufCommand demonstrates Protobuf data modeling and binary serialization.
var ProtobufCommand *cli.Command = &cli.Command{
	Name:        "protobuf",
	Usage:       "Demonstrate Protobuf data modeling",
	Description: "This command round-trips a Profile through the Protobuf wire format.",
	Action: func(context.Context, *cli.Command) error {
		profile, wireData, err := practical.ProfileRoundTrip()
		if err != nil {
			return err
		}

		profileJSON, err := (protojson.MarshalOptions{Multiline: true, Indent: "  "}).Marshal(profile)
		if err != nil {
			return err
		}

		fmt.Printf("Profile after binary round-trip:\n%s\n", profileJSON)
		fmt.Printf("Wire data (hex): %s\n", hex.EncodeToString(wireData))
		return nil
	},
}

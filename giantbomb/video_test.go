package giantbomb

import (
	"fmt"
	"os"
	"testing"
)

func TestVideo(t *testing.T) {

	client := NewClient(os.Getenv("GB_API_KEY"))

	listParams := ListVideoRequest{
		Limit:  "5",
		Filter: "name:pure gameplay montage",
	}

	videos, err := client.Video.List(&listParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	assertEqual(t, len(videos.Results), 1, "Expected Single Result")
	assertEqual(t, videos.Results[0].Name, "Pure Gameplay Montage", "Expected Pure Gameplay Montage")

	guid := videos.Results[0].GUID

	video, err := client.Video.Get(guid, nil)

	assertEqual(t, video.Results.GUID, guid, "Unexpected GUID")
	assertEqual(t, video.Results.Name, "Pure Gameplay Montage", "Expected Pure Gameplay Montage")
}

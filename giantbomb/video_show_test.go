package giantbomb

import (
	"fmt"
	"os"
	"testing"
)

func TestVideoShow(t *testing.T) {

	client := NewClient(os.Getenv("GB_API_KEY"))

	listParams := ListVideoShowRequest{
		Limit:  "1",
		Filter: "guid:2340-2",
	}

	videoShows, err := client.VideoShow.List(&listParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	assertEqual(t, len(videoShows.Results), 1, "Expected Single Result")
	assertEqual(t, videoShows.Results[0].Title, "Endurance Run: Persona 4", "Expected Endurance Run: Persona 4")

	guid := videoShows.Results[0].GUID

	videoShow, err := client.VideoShow.Get(guid, nil)

	assertEqual(t, videoShow.Results.GUID, guid, "Unexpected GUID")
	assertEqual(t, videoShow.Results.Title, "Endurance Run: Persona 4", "Expected Endurance Run: Persona 4")
}

package giantbomb

import (
	"fmt"
	"os"
	"testing"
)

func TestGame(t *testing.T) {

	client := NewClient(os.Getenv("GB_API_KEY"))

	listParams := ListGamesRequest{
		Limit:  "5",
		Filter: "name:killer7",
	}

	games, err := client.Game.List(&listParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	assertEqual(t, len(games.Results), 1, "Expected Single Result")
	assertEqual(t, games.Results[0].Name, "Killer7", "Expected Killer7")

	guid := games.Results[0].GUID

	game, err := client.Game.Get(guid, nil)

	assertEqual(t, game.Results.GUID, guid, "Unexpected GUID")
	assertEqual(t, game.Results.Name, "Killer7", "Expected Killer7")
}

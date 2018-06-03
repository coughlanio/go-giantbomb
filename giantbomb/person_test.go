package giantbomb

import (
	"fmt"
	"os"
	"testing"
)

func TestPerson(t *testing.T) {

	client := NewClient(os.Getenv("GB_API_KEY"))

	listParams := ListPersonRequest{
		Limit:  "5",
		Filter: "name:peter molyneux",
	}

	persons, err := client.Person.List(&listParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	assertEqual(t, len(persons.Results), 1, "Expected Single Result")
	assertEqual(t, persons.Results[0].Name, "Peter Molyneux", "Expected Peter Molyneux")

	guid := persons.Results[0].GUID

	person, err := client.Person.Get(guid, nil)

	assertEqual(t, person.Results.GUID, guid, "Unexpected GUID")
	assertEqual(t, person.Results.Name, "Peter Molyneux", "Expected Peter Molyneux")
}

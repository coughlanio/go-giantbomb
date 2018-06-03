package giantbomb

import (
	"fmt"
)

// PersonService handles communication with the person/people endpoints on GiantBomb.
type PersonService service

// GetPersonRequest is the standard request struct for the person endpoint on Giant Bomb.
type GetPersonRequest struct {
	FieldList string `json:"field_list"`
}

// ListPersonRequest is the standard request struct for the people endpoint on Giant Bomb.
type ListPersonRequest struct {
	FieldList      string `json:"field_list"`
	Limit          string `json:"limit"`
	Offset         string `json:"offset"`
	SubscriberOnly string `json:"subscriber_only"`
	Sort           string `json:"sort"`
	Filter         string `json:"filter"`
}

// GetPersonResponse is the standard response from the person endpoint on Giant Bomb.
type GetPersonResponse struct {
	*Response
	Results Person `json:"results"`
}

// ListPersonResponse is the standard response from the people endpoint on Giant Bomb.
type ListPersonResponse struct {
	*Response
	Results []Person `json:"results"`
}

// Person is the base API person details struct.
type Person struct {
	Aliases           string     `json:"aliases"`
	APIDetailURL      string     `json:"api_detail_url"`
	BirthDate         string     `json:"birth_date"`
	Characters        []Detail   `json:"characters"`
	Concepts          []Detail   `json:"concepts"`
	Country           string     `json:"country"`
	DateAdded         string     `json:"date_added"`
	DateLastUpdated   string     `json:"date_last_updated"`
	DeathDate         string     `json:"death_date"`
	Deck              string     `json:"deck"`
	Description       string     `json:"description"`
	FirstCreditedGame Detail     `json:"first_credited_game"`
	Franchises        []Detail   `json:"franchises"`
	Games             []Detail   `json:"games"`
	Gender            int        `json:"gender"`
	GUID              string     `json:"guid"`
	Hometown          string     `json:"hometown"`
	ID                int        `json:"id"`
	Image             Image      `json:"image"`
	ImageTags         []ImageTag `json:"image_tags"`
	Locations         []Detail   `json:"locations"`
	Name              string     `json:"name"`
	Objects           []Detail   `json:"objects"`
	People            []Detail   `json:"people"`
	SiteDetailURL     string     `json:"site_detail_url"`
}

// Get returns an instance of a person based on the guid.
func (s *PersonService) Get(guid string, params *GetPersonRequest) (*GetPersonResponse, error) {

	path := fmt.Sprintf("person/%s", guid)

	req, err := s.client.NewRequest("GET", path, params)

	if err != nil {
		return nil, err
	}

	var person *GetPersonResponse

	s.client.Do(req, &person)

	if err != nil {
		return nil, err
	}

	return person, nil
}

// List returns a paginated list of all persons.
func (s *PersonService) List(params *ListPersonRequest) (*ListPersonResponse, error) {

	req, err := s.client.NewRequest("GET", "people", params)

	if err != nil {
		return nil, err
	}

	var people *ListPersonResponse

	s.client.Do(req, &people)

	if err != nil {
		return nil, err
	}

	return people, nil
}

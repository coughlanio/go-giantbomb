package giantbomb

import (
	"fmt"
)

// GameService handles communication with the games endpoints on GiantBomb.
type GameService service

// GetGameRequest is the standard request struct for the game endpoint on Giant Bomb.
type GetGameRequest struct {
	FieldList string `json:"field_list"`
}

// ListGamesRequest is the standard request struct for the games endpoint on Giant Bomb.
type ListGamesRequest struct {
	FieldList string `json:"field_list"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Platforms string `json:"platforms"`
	Sort      string `json:"sort"`
	Filter    string `json:"filter"`
}

// GetGameResponse is the standard response from the game endpoint on Giant Bomb.
type GetGameResponse struct {
	*Response
	Results Game `json:"results"`
}

// ListGamesResponse is the standard response from the games endpoint on Gian Bomb.
type ListGamesResponse struct {
	*Response
	Results []Games `json:"results"`
}

// Game is the base API game details struct.
type Game struct {
	Aliases                string      `json:"aliases"`
	APIDetailURL           string      `json:"api_detail_url"`
	DateAdded              string      `json:"date_added"`
	DateLastUpdated        string      `json:"date_last_updated"`
	Deck                   string      `json:"deck"`
	Description            string      `json:"description"`
	ExpectedReleaseDay     string      `json:"expected_release_day"`
	ExpectedReleaseMonth   string      `json:"expected_release_month"`
	ExpectedReleaseQuarter string      `json:"expected_release_quarter"`
	ExpectedReleaseYear    string      `json:"expected_release_year"`
	GUID                   string      `json:"guid"`
	ID                     int         `json:"id"`
	Image                  Image       `json:"image"`
	ImageTags              []ImageTag  `json:"image_tags"`
	Name                   string      `json:"name"`
	NumberOfUserReviews    int         `json:"number_of_user_reviews"`
	OriginalGameRating     []Detail    `json:"original_game_rating"`
	OriginalReleaseDate    string      `json:"original_release_date"`
	Platforms              []Platforms `json:"platforms"`
	SiteDetailURL          string      `json:"site_detail_url"`
}

// Games is the extended API game details struct.
type Games struct {
	*Game
	Videos                    []Detail `json:"videos"`
	Characters                []Detail `json:"characters"`
	Concepts                  []Detail `json:"concepts"`
	Developers                []Detail `json:"developers"`
	FirstAppearanceCharacters []Detail `json:"first_appearance_characters"`
	FirstAppearanceConcepts   []Detail `json:"first_appearance_concepts"`
	FirstAppearanceLocations  []Detail `json:"first_appearance_locations"`
	FirstAppearanceObjects    []Detail `json:"first_appearance_objects"`
	FirstAppearancePeople     []Detail `json:"first_appearance_people"`
	Franchises                []Detail `json:"franchises"`
	Genres                    []Detail `json:"genres"`
	KilledCharacters          []Detail `json:"killed_characters"`
	Locations                 []Detail `json:"locations"`
	Objects                   []Detail `json:"objects"`
	People                    []Detail `json:"people"`
	Publishers                []Detail `json:"publishers"`
	Releases                  []Detail `json:"releases"`
	SimilarGames              []Detail `json:"similar_games"`
	Themes                    []Detail `json:"themes"`
}

// Platforms extends the base API link details structure with support for platform abbreviations.
type Platforms struct {
	*Detail
	Abbreviation string `json:"abbreviation"`
}

// Get queries a game by GUID.
func (s *GameService) Get(guid string, params *GetGameRequest) (*GetGameResponse, error) {

	path := fmt.Sprintf("game/%s", guid)

	req, err := s.client.NewRequest("GET", path, params)

	if err != nil {
		return nil, err
	}

	var game *GetGameResponse

	s.client.Do(req, &game)

	if err != nil {
		return nil, err
	}

	return game, nil
}

// List retrieves a paginated list of all games on Giant Bomb.
func (s *GameService) List(params *ListGamesRequest) (*ListGamesResponse, error) {

	req, err := s.client.NewRequest("GET", "games", params)

	if err != nil {
		return nil, err
	}

	var games *ListGamesResponse

	s.client.Do(req, &games)

	if err != nil {
		return nil, err
	}

	return games, nil
}

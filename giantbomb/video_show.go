package giantbomb

import "fmt"

// VideoShowService handles communication with the video shows endpoints on GiantBomb.
type VideoShowService service

// GetVideoShowRequest is the standard request struct for the video show endpoint on Giant Bomb.
type GetVideoShowRequest struct {
	FieldList string `json:"field_list"`
}

// ListVideoShowRequest is the standard request struct for the video shows endpoint on Giant Bomb.
type ListVideoShowRequest struct {
	FieldList string `json:"field_list"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Sort      string `json:"sort"`
	Filter    string `json:"filter"`
}

// GetVideoShowResponse is the response object for Get
type GetVideoShowResponse struct {
	*Response
	Results VideoShow `json:"results"`
}

// ListVideoShowResponse is the response object for List
type ListVideoShowResponse struct {
	*Response
	Results []VideoShow `json:"results"`
}

// VideoShow is the base API video show details struct.
type VideoShow struct {
	APIDetailURL  string  `json:"api_detail_url"`
	SiteDetailURL string  `json:"site_detail_url"`
	Deck          string  `json:"deck"`
	GUID          string  `json:"guid"`
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Position      int     `json:"position"`
	DisplayNav    bool    `json:"display_nav"`
	Active        bool    `json:"active"`
	Image         Image   `json:"image"`
	Latest        []Video `json:"latest"`
}

// Get returns a instance of a video show based on guid.
func (s *VideoShowService) Get(guid string, params *GetVideoShowRequest) (*GetVideoShowResponse, error) {

	path := fmt.Sprintf("video_show/%s", guid)

	req, err := s.client.NewRequest("GET", path, params)

	if err != nil {
		return nil, err
	}

	var videoShow *GetVideoShowResponse

	s.client.Do(req, &videoShow)

	if err != nil {
		return nil, err
	}

	return videoShow, nil
}

// List returns a paginated list of all video shows.
func (s *VideoShowService) List(params *ListVideoShowRequest) (*ListVideoShowResponse, error) {

	req, err := s.client.NewRequest("GET", "video_shows", params)

	if err != nil {
		return nil, err
	}

	var videoShows *ListVideoShowResponse

	s.client.Do(req, &videoShows)

	if err != nil {
		return nil, err
	}

	return videoShows, nil
}

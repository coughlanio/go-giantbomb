package giantbomb

import "fmt"

// VideoService handles communication with the videos endpoints on GiantBomb.
type VideoService service

// GetVideoRequest is the standard request struct for the video endpoint on Giant Bomb.
type GetVideoRequest struct {
	FieldList string `json:"field_list"`
}

// ListVideoRequest is the standard request struct for the videos endpoint on Giant Bomb.
type ListVideoRequest struct {
	FieldList      string `json:"field_list"`
	Limit          string `json:"limit"`
	Offset         string `json:"offset"`
	SubscriberOnly string `json:"subscriber_only"`
	Sort           string `json:"sort"`
	Filter         string `json:"filter"`
}

// GetVideoResponse is the standard response from the video endpoint on Giant Bomb.
type GetVideoResponse struct {
	*Response
	Results Video `json:"results"`
}

// ListVideoResponse is the standard response from the videos endpoint on Giant Bomb.
type ListVideoResponse struct {
	*Response
	Results []Video `json:"results"`
}

// Video is the base API video details struct.
type Video struct {
	APIDetailURL    string    `json:"api_detail_url"`
	Deck            string    `json:"deck"`
	HdURL           string    `json:"hd_url"`
	HighURL         string    `json:"high_url"`
	LowURL          string    `json:"low_url"`
	EmbedPlayer     string    `json:"embed_player"`
	GUID            string    `json:"guid"`
	ID              int       `json:"id"`
	LengthSeconds   int       `json:"length_seconds"`
	Name            string    `json:"name"`
	PublishDate     string    `json:"publish_date"`
	SiteDetailURL   string    `json:"site_detail_url"`
	URL             string    `json:"url"`
	Image           Image     `json:"image"`
	User            string    `json:"user"`
	VideoType       string    `json:"video_type"`
	VideoShow       VideoShow `json:"video_show"`
	VideoCategories []Detail  `json:"video_categories"`
	SavedTime       string    `json:"saved_time"`
	YoutubeID       string    `json:"youtube_id"`
}

// Get returns an instance of a video based on the guid.
func (s *VideoService) Get(guid string, params *GetVideoRequest) (*GetVideoResponse, error) {

	path := fmt.Sprintf("video/%s", guid)

	req, err := s.client.NewRequest("GET", path, params)

	if err != nil {
		return nil, err
	}

	var video *GetVideoResponse

	s.client.Do(req, &video)

	if err != nil {
		return nil, err
	}

	return video, nil
}

// List returns a paginated list of all videos.
func (s *VideoService) List(params *ListVideoRequest) (*ListVideoResponse, error) {

	req, err := s.client.NewRequest("GET", "videos", params)

	if err != nil {
		return nil, err
	}

	var videos *ListVideoResponse

	s.client.Do(req, &videos)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

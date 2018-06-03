package giantbomb

// SavedTimeService handles communication with the saved-times endpoints on GiantBomb.
type SavedTimeService service

// GetSavedTimeRequest is the standard request object for the get-saved-time endpoint on Giant Bomb.
type GetSavedTimeRequest struct {
	VideoID int `json:"video_id,string"`
}

// GetSavedTimeResponse is the standard response from the get-saved-time endpoint on Giant Bomb.
type GetSavedTimeResponse struct {
	Success   int    `json:"success"`
	SavedTime string `json:"savedTime"`
}

// ListSavedTimeResponse is the standard response from the get-all-saved-times endpoint on Giant Bomb.
type ListSavedTimeResponse struct {
	Success    int `json:"success"`
	SavedTimes []struct {
		VideoID   int    `json:"videoId"`
		SavedTime string `json:"savedTime"`
		SavedOn   string `json:"savedOn"`
	} `json:"savedTimes"`
}

// Get retrieves a saved time for a specific video_id on Giant Bomb.
func (s *SavedTimeService) Get(params *GetSavedTimeRequest) (*GetSavedTimeResponse, error) {

	req, err := s.client.NewRequest("GET", "video/get-saved-time", params)

	if err != nil {
		return nil, err
	}

	var savedTime *GetSavedTimeResponse

	s.client.Do(req, &savedTime)

	if err != nil {
		return nil, err
	}

	return savedTime, nil
}

// List retrieves a list of all saved times on Giant Bomb.
func (s *SavedTimeService) List() (*ListSavedTimeResponse, error) {

	req, err := s.client.NewRequest("GET", "video/get-all-saved-times", nil)

	if err != nil {
		return nil, err
	}

	var savedTimes *ListSavedTimeResponse

	s.client.Do(req, &savedTimes)

	if err != nil {
		return nil, err
	}

	return savedTimes, nil
}

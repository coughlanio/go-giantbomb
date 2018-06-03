package giantbomb

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://www.giantbomb.com/api/"
	defaultUserAgent = "go-giantbomb/0.1 (+https://github.com.com/coughlanio/go-giantbomb)"
)

// Params is the generic params struct.
type Params struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

// Client manages communication with the Giant Bomb API.
type Client struct {
	client *http.Client

	BaseURL   *url.URL
	APIKey    string
	UserAgent string

	common service

	Game      *GameService
	Video     *VideoService
	VideoShow *VideoShowService
	Person    *PersonService
	SavedTime *SavedTimeService
}

// Service is the generic service that wraps the client.
type service struct {
	client *Client
}

// Response is the generic API response structure for the Giant Bomb API.
type Response struct {
	Error                string `json:"error"`
	Limit                int    `json:"limit"`
	Offset               int    `json:"offset"`
	NumberOfPageResults  int    `json:"number_of_page_results"`
	NumberOfTotalResults int    `json:"number_of_total_results"`
	StatusCode           int    `json:"status_code"`
	Version              string `json:"version"`
}

// Detail is the base API link details data structure.
type Detail struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

// ImageTag extends the base API link details structure with support for tag counts.
type ImageTag struct {
	*Detail
	Total int `json:"total"`
}

// Image is the base API image details struct.
type Image struct {
	IconURL        string `json:"icon_url"`
	MediumURL      string `json:"medium_url"`
	ScreenURL      string `json:"screen_url"`
	ScreenLargeURL string `json:"screen_large_url"`
	SmallURL       string `json:"small_url"`
	SuperURL       string `json:"super_url"`
	ThumbURL       string `json:"thumb_url"`
	TinyURL        string `json:"tiny_url"`
	OriginalURL    string `json:"original_url" json:"original"`
	ImageTags      string `json:"image_tags" json:"tags"`
}

// NewClient returns a new instance of the API client.
func NewClient(apiKey string) *Client {

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    http.DefaultClient,
		BaseURL:   baseURL,
		APIKey:    apiKey,
		UserAgent: defaultUserAgent,
	}

	c.common.client = c

	c.Game = (*GameService)(&c.common)
	c.Video = (*VideoService)(&c.common)
	c.VideoShow = (*VideoShowService)(&c.common)
	c.Person = (*PersonService)(&c.common)
	c.SavedTime = (*SavedTimeService)(&c.common)

	return c
}

// NewRequest assembles a new request to be executed.
func (c *Client) NewRequest(method string, path string, params interface{}) (*http.Request, error) {

	u, _ := c.BaseURL.Parse(path)

	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	q := req.URL.Query()

	q.Add("api_key", c.APIKey)
	q.Add("format", "json")

	p := ConvertToMap(params)

	for key, value := range p {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Do executes the request against the Giant Bomb API.
func (c *Client) Do(req *http.Request, v interface{}) error {

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)

	decErr := json.Unmarshal(buff, &v)

	if decErr == io.EOF {
		decErr = nil
	}
	if decErr != nil {
		err = decErr
	}

	return err
}

// ConvertToMap converts a generic interface to a string map.
func ConvertToMap(params interface{}) map[string]string {

	var i map[string]string

	if params == nil {
		return i
	}

	j, _ := json.Marshal(params)

	err := json.Unmarshal([]byte(j), &i)

	if err != nil {
		fmt.Println(err)
	}

	return i
}

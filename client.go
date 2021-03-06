package octopus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"github.com/pkg/errors"
)

const ApiBaseUrl = "https://strong-octopus.com"

type Client struct {
	presharedKey string
	httpClient http.Client
}

func NewClient(presharedKey string) Client {
	return Client{
		presharedKey: presharedKey,
		httpClient: http.Client{},
	}
}

type ApiError struct {
	Message string
}
func (e *ApiError) Error() string {
	return e.Message
}

func (c *Client) SearchByKeyword(keyword string, page int) ([]Article, error) {
	const path = "/articles/search"

	params := url.Values{}
	params.Add("keyword", keyword)
	params.Add("page", strconv.Itoa(page))
	responseBody, err := c.request("GET", path, params)
	if err != nil {
		return nil, err
	}

	var searchResponse struct {
		Articles []Article `json:"articles"`
	}

	if err := json.Unmarshal(responseBody, &searchResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode json")
	}

	return searchResponse.Articles, nil
}

func (c *Client) request(method string, path string, params url.Values) ([]byte, error) {
	var endpoint = fmt.Sprintf("%s%s?%s", ApiBaseUrl, path, params.Encode())

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.presharedKey))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var errorResponse struct {
			Message string `json:"error"`
		}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return nil, errors.Wrap(err, "failed to decode json")
		}

		return nil, &ApiError{Message: errorResponse.Message}
	}

	return body, nil
}

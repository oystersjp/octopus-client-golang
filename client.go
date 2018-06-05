package octopus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ApiBaseUrl = "https://strong-octopus.com"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}

func (c *Client) SearchByKeyword(keyword string, page int) ([]Article, error) {
	var path = fmt.Sprintf("/searchResponse/search?keyword=%s&page=%d", keyword, page)
	var endpoint = fmt.Sprintf("%s%s", ApiBaseUrl, path;

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

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
			return nil, err
		}

		return nil, errors.New(errorResponse.Message)
	}

	var searchResponse struct {
		Articles []Article `json:"searchResponse"`
	}

	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, err
	}

	return searchResponse.Articles, nil
}

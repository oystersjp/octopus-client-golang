package octopus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://strong-octopus.com/articles/search?keyword=%s&page=%d"

func SearchByKeyword(keyword string, page int) ([]Article, error) {

	url := fmt.Sprintf(apiUrl, keyword, page)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorMsg ErrorMessage
		if err := json.Unmarshal(body, &errorMsg); err != nil {
			return nil, err
		}

		err := errors.New(errorMsg.Message)
		return nil, err
	}

	var articles Articles
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}

	return articles.Articles, nil
}

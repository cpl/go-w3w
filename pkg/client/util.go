package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cpl/go-w3w/pkg/w3w"
)

const HeaderAPIKey = "X-Api-Key"

func checkResponse(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		var errResponse w3w.ResponseError

		if err := json.NewDecoder(res.Body).Decode(&errResponse); err != nil {
			return fmt.Errorf(
				"bad response status code, %d, failed to parse response, %w",
				res.StatusCode, err)
		}

		return fmt.Errorf(
			"bad response status code, %d, %s - %s",
			res.StatusCode, errResponse.Error.Code, errResponse.Error.Message)
	}

	return nil
}

func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.Endpoint+url, body)
	if err != nil {
		return nil, fmt.Errorf("failed creating new client request, %w", err)
	}

	req.URL.Query().Add("language", c.Language)
	req.Header.Set(HeaderAPIKey, c.key)

	return req, nil
}

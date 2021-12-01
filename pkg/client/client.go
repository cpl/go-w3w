package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/eljoth/go-w3w/pkg/w3w"
)

const EndpointDefault = "https://api.what3words.com/v3/"

type Client struct {
	Endpoint string
	key      string
}

func New(key string) *Client {
	return &Client{
		key:      key,
		Endpoint: EndpointDefault,
	}
}

func (c *Client) ConvertToCoordinates(address string) (*w3w.ResponseConvert, error) {
	address=url.QueryEscape(address)
	req, err := c.newRequest(
		http.MethodGet,
		"convert-to-coordinates?words="+address, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed making request, %w", err)
	}
	defer res.Body.Close()

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	response := new(w3w.ResponseConvert)
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, fmt.Errorf("failed decoding response, %w", err)
	}

	return response, nil
}

func (c *Client) ConvertTo3WordAddress(coordinates w3w.Coordinates) (*w3w.ResponseConvert, error) {
	req, err := c.newRequest(
		http.MethodGet,
		"convert-to-3wa?coordinates="+coordinates.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed making request, %w", err)
	}
	defer res.Body.Close()

	if err := checkResponse(res); err != nil {
		return nil, err
	}

	response := new(w3w.ResponseConvert)
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, fmt.Errorf("failed decoding response, %w", err)
	}

	return response, nil
}

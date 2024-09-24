package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type ResponseBody []byte
type Headers map[string]string
type QueryParams map[string]string

type Client struct {
	client      *http.Client
	headers     Headers
	method      string
	url         string
	payload     io.Reader
	StatusCode  int
	log         *slog.Logger
	queryParams QueryParams
}

func New(baseURL string) *Client {
	return &Client{
		// this is baseUrl only, need to updated for specific path later using other func
		url:     baseURL,
		client:  &http.Client{},
		headers: make(map[string]string),
	}
}

func (c *Client) Token(token string) *Client {
	c.headers["Authorization"] = "Bearer " + token
	return c
}

func (c *Client) GetJSON(path string, response interface{}, query map[string]string) (int, error) {
	c.headers["Content-Type"] = "application/json"
	c.method = "GET"
	c.url = c.url + path + c.convertMapToQuery(query)

	return c.sendRequest(response)
}

func (c *Client) sendRequest(responseStruct interface{}) (int, error) {
	req, err := http.NewRequest(c.method, c.url, c.payload)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	c.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("error reading response body: %v", err)
	}

	// no need to unmarshall the body, since there is no content
	if c.StatusCode != http.StatusNoContent {
		if err := c.unmarshalJSON(body, responseStruct); err != nil {
			return resp.StatusCode, fmt.Errorf("error unmarshaling response body: %v", err)
		}
	}

	return resp.StatusCode, nil
}

func (c *Client) unmarshalJSON(responseBody ResponseBody, target interface{}) error {
	err := json.Unmarshal(responseBody, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func (c *Client) convertMapToQuery(params map[string]string) string {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}

	return query.Encode()
}

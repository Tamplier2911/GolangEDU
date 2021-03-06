package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type HTTPClient struct {
	Client *http.Client
}

// Initialize and setup HTTPCleint, returns instance of HTTPCleint
func (c HTTPClient) Setup() HTTPClient {
	// init client
	c.Client = &http.Client{}
	return c
}

// Initialize GET request using underlying client, unmarshal data to provided data type
func (c *HTTPClient) Get(url string, i interface{}) error {
	// init api call
	log.Println("request", url)
	resp, err := c.Client.Get(url)
	if err != nil {
		return errors.New("request failed")
	}
	log.Println("response", resp)
	defer resp.Body.Close()

	// read res body
	log.Println("reading responce body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}
	// log.Println("body", string(body[:]))

	// unmarshal res body
	log.Println("unmarshaling response body")
	err = json.Unmarshal(body, &i)
	if err != nil {
		return errors.New("failed to unmarshal response body")
	}

	return nil
}

// Iitialize GET request using underlying client, returns raw bytes
func (c *HTTPClient) GetRaw(url string) ([]byte, error) {
	// init api call
	log.Println("request", url)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, errors.New("request failed")
	}
	// log.Println("response", resp)
	defer resp.Body.Close()

	// read res body
	log.Println("reading responce body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	return body, nil
}

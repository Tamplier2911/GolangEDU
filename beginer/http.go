package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Initialize and setup HTTPCleint, returns instance of HTTPCleint
func (c HTTPClient) Setup() HTTPClient {
	// init client
	c.Client = &http.Client{}
	return c
}

// Initialize GET request using underlying client
func (c *HTTPClient) Get(url string, i interface{}) error {
	// init api call
	log.Println("request", url)
	resp, err := c.Client.Get(url)
	if err != nil {
		log.Fatal("request failed", err)
		return errors.New("request failed")
	}
	log.Println("response", resp)
	defer resp.Body.Close()

	// read res body
	log.Println("reading responce body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed to read response body", err)
		return errors.New("failed to read response body")
	}

	// unmarshal res body
	log.Println("unmarshaling response body")
	err = json.Unmarshal(body, &i)
	if err != nil {
		log.Fatal("failed to unmarshal response body", err)
		return errors.New("failed to unmarshal body")
	}

	return nil
}

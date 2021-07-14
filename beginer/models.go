package main

import "net/http"

type HTTPClient struct {
	Client *http.Client
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

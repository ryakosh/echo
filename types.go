package main

import (
	"net/http"
	"net/url"
)

// Echo stores information about client's request and server's response
type Echo struct {
	Version  string   `json:"version"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

// Request stores information about request's method, url and headers
type Request struct {
	Method  string      `json:"method"`
	URL     URL         `json:"url"`
	Headers http.Header `json:"headers"`
	Body    Body        `json:"body"`
}

// Response stores information about response's status and headers
type Response struct {
	Status  Status      `json:"status"`
	Headers http.Header `json:"headers"`
}

// URL stores information about url's path and query-string
type URL struct {
	Raw   string `json:"raw"`
	Path  string `json:"path"`
	Query Query  `json:"query"`
}

// Query stores information about query-string, including it's parsed form
type Query struct {
	Raw    string     `json:"raw"`
	Parsed url.Values `json:"parsed"`
}

// Body stores information about request's body, including it's parsed form
type Body struct {
	Raw    string                 `json:"raw"`
	Parsed map[string]interface{} `json:"parsed"`
}

// Status stores information about response's status code and text
type Status struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

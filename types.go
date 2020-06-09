package main

import (
	"net/http"
	"net/url"
)

type Echo struct {
	Version  string   `json:"version"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	Method  string      `json:"method"`
	URL     URL         `json:"url"`
	Headers http.Header `json:"headers"`
}

type Response struct {
	Status  Status      `json:"status"`
	Headers http.Header `json:"headers"`
}

type URL struct {
	Raw   string `json:"raw"`
	Path  string `json:"path"`
	Query Query  `json:"query"`
}

type Query struct {
	Raw    string     `json:"raw"`
	Parsed url.Values `json:"parsed"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

package gohttp

import (
	"io"
	"net/http"
)

type Handler interface {
	OnRequestStart(*http.Request)
	OnRequestComplete(*http.Request, *http.Response)
	OnError(*http.Request, error)
}

type Worker interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Put(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Patch(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
	AddHandlers(handlers ...Handler)
}

type ConfigFunc func(*Client)

type Client struct {
	client *http.Client
}

func NewClient(config ...ConfigFunc) Worker {
	c := &Client{}

	for _, cFunc := range config {
		cFunc(c)
	}
	if c.client == nil {
		c.client = &http.Client{}
	}

	return c
}

func (c *Client) do(method, url string, headers http.Header, body io.Reader) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Get(url string, headers http.Header) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Post(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Put(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Patch(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Delete(url string, headers http.Header) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	panic("implement me")
}

func (c *Client) AddHandlers(handlers ...Handler) {
	panic("implement me")
}

var _ Worker = (*Client)(nil)

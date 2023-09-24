package yy

import (
	"fmt"
	"time"
)

type Options struct {
	timeout time.Duration
	retry uint
}

type Option func(*Options)

func Timeout(timeout time.Duration) Option {
	return func (options *Options)  {
		options.timeout = timeout
	}
}

func Retry(retry uint) Option {
	return func(options *Options) {
		options.retry = retry
	}
}

func NewClient(host string, port uint16, setters ...Option) *Client {
	options := &Options {
		timeout: 1 * time.Second,
		retry: 2,
	}
	for _, setter := range setters {
		setter(options)
	}
	return &Client{
		host: host,
		port: port,
		timeout: options.timeout,
		retry: options.retry,
	}
}

func (c *Client) String() string {
	return fmt.Sprint(*c)
}

type Client struct {
	host string
	port uint16
	timeout time.Duration
	retry uint
}
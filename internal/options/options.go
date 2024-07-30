package options

import (
	"net/http"
	"net/url"
)

type Options struct {
	Headers http.Header
	URL     *url.URL
	Method  string
}

type Option func(*Options)

func WithHeader(k, v string) Option {
	return func(opts *Options) {
		opts.Headers.Set(k, v)
	}
}

func WithURL(u *url.URL) Option {
	return func(opts *Options) {
		opts.URL = u
	}
}

func WithMethod(m string) Option {
	return func(opts *Options) {
		opts.Method = m
	}
}

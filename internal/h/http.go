package h

import (
	"bytes"
	"net/http"

	"go.mdl.wtf/zhw/internal/options"
)

type Client struct {
	*http.Client
	Options *options.Options
}

func (c *Client) Send(d []byte) error {
	req, err := http.NewRequest(c.Options.Method, c.Options.URL.String(), bytes.NewReader(d))
	if err != nil {
		return err
	}
	req.Header = c.Options.Headers
	_, err = c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func New(opts *options.Options) (*Client, error) {
	d := http.DefaultClient
	c := &Client{d, opts}
	return c, nil
}

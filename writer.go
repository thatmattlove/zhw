package zhw

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"go.mdl.wtf/zhw/internal/h"
	"go.mdl.wtf/zhw/internal/options"
)

type (
	Option  = options.Option
	Options = options.Options
)

var (
	WithHeader = options.WithHeader
	WithURL    = options.WithURL
	WithMethod = options.WithMethod
)

const (
	delim string = ","
)

type Writer struct {
	client *h.Client
	w      *bytes.Buffer
	bw     *bufio.Writer
	once   *sync.Once
}

func (w *Writer) Write(b []byte) (int, error) {
	n, err := w.bw.Write(b)
	if err != nil {
		return 0, err
	}
	_, err = w.bw.Write([]byte(delim))
	if err != nil {
		return 0, err
	}
	return n + 1, nil
}

// Close flushes the log buffer and sends all buffered logs as a JSON array to the configured HTTP host.
func (w *Writer) Close() error {
	w.once.Do(func() {
		err := w.bw.Flush()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		// Remove last comma
		w.w.Truncate(w.w.Len() - 1)
		// Close array
		_, err = w.w.Write([]byte(`]`))
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		b := new(bytes.Buffer)
		err = json.Compact(b, w.w.Bytes())
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		err = w.client.Send(b.Bytes())
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
	})
	return nil
}

// NewWriter creates a new io.Writer (io.WriteCloser, actually) that buffers log messages from
// zerolog and sends all buffered messages as a JSON array when Close() is called.
func NewWriter(o ...options.Option) (*Writer, error) {
	opts := &options.Options{
		Method:  http.MethodPost,
		Headers: http.Header{"content-type": []string{"application-json"}},
	}
	for _, opt := range o {
		opt(opts)
	}
	if opts.URL == nil {
		return nil, fmt.Errorf("URL option is required")
	}
	client, err := h.New(opts)
	if err != nil {
		return nil, err
	}
	w := new(bytes.Buffer)
	bw := bufio.NewWriter(w)

	// Open array
	bw.Write([]byte(`[`))

	wr := &Writer{
		client: client,
		w:      w,
		bw:     bw,
		once:   new(sync.Once),
	}
	return wr, nil
}

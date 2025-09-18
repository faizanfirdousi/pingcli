//This file defines a small package httpclient that wraps net/http to
// provide a configurable HTTP client and a Get(ctx, url) method that
// measures request time and returns a compact Response struct.

package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Client interface for HTTP operations (enables mocking)
type Client interface {
	//Method Signature
	Get(ctx context.Context, url string) (*Response, error)
}

//Response Wraps HTTP response with timing data

type Response struct {
	StatusCode    int
	Status        string
	Duration      time.Duration
	ContentLength int64
}

//HTTPClient implements the Client interface
type HTTPClient struct {
	client *http.Client
	timeout time.Duration
}

//Config for HTTP client setup
type Config struct {
	Timeout         time.Duration
	FollowRedirects bool
	MaxRedirects    int
	UserAgent       string
}

//NewHTTPClient creates a new HTTP cleint with configuration
func NewHTTPClient(config Config) *HTTPClient{
	transport := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: false,
	}

	client := &http.Client{
		Transport: transport,
		Timeout: config.Timeout,
	}

	//Configure redirects
	if !config.FollowRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error{
			return http.ErrUseLastResponse
		}
	} else if config.MaxRedirects > 0 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= config.MaxRedirects {
				return fmt.Errorf("maximum redirects exceeded: %d", config.MaxRedirects)
			}
			return nil
		}
	}

	return &HTTPClient{
		client: client,
		timeout: config.Timeout,
	}
}

//Get performs HTTP GET request with timing

func (c *HTTPClient) Get(ctx context.Context, url string) (*Response, error){
	//Validate URL
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w",err)
	}

	//Set headers
	req.Header.Set("User-Agent", "PingCLI/1.0")
	req.Header.Set("Accept","*/*")

	start := time.Now()
	resp, err := c.client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w",err)
	}

	defer resp.Body.Close()

	return &Response{
		StatusCode: resp.StatusCode,
		Status: resp.Status,
		Duration: duration,
		ContentLength: resp.ContentLength,
	}, nil


}

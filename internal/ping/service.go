package ping

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/faizanfirdousi/pingcli/pkg/httpclient"
)

//PingResult reporesents the result of a ping operation

type PingResult struct{
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
	Status string `json:"status"`
	Duration time.Duration `json:"duration"`
	Succeess bool `json:"success"`
	Error string `json:"error,omitempty"`
}

type Service struct {
	client httpclient.Client
}

func NewService(client httpclient.Client) *Service{
	return &Service{
		client: client,
	}
}

// Ping executes a single ping connection
func (s *Service) Ping(ctx context.Context, targetURL string) *PingResult{
	result := &PingResult{
		URL: targetURL,
	}

	//Validate URL format
	if err := s.validateURL(targetURL); err != nil {
		result.Error = fmt.Sprintf("invalid URL: %v", err)
		return result
	}
}

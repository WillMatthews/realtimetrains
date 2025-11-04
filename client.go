package realtimetrains

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// DefaultBaseURL is the production RTT API base. It ends without a trailing slash.
const DefaultBaseURL = "https://api.rtt.io/api/v1"

// Client wraps access to the RealTimeTrains Pull ("simple") JSON API.
// It is safe for concurrent use.
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	user       string
	pass       string
	// optional: per-request user agent
	userAgent string
}

// Option allows customizing Client.
type Option func(*Client) error

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) error {
		if hc != nil {
			c.httpClient = hc
		}
		return nil
	}
}

// WithBaseURL overrides the API base. Accepts values with or without trailing slash.
func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		if raw == "" {
			return errors.New("base url empty")
		}
		u, err := url.Parse(strings.TrimRight(raw, "/"))
		if err != nil {
			return err
		}
		c.baseURL = u
		return nil
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(ua string) Option { return func(c *Client) error { c.userAgent = ua; return nil } }

// New creates a new API client. Username & password are the RTT credentials (Basic Auth).
func New(username, password string, opts ...Option) (*Client, error) {
	if username == "" || password == "" {
		return nil, errors.New("username & password required")
	}
	base, _ := url.Parse(DefaultBaseURL)
	c := &Client{
		httpClient: &http.Client{Timeout: 15 * time.Second},
		baseURL:    base,
		user:       username,
		pass:       password,
		userAgent:  "gorealtimetrains/0.1 (+https://github.com/WillMatthews/gorealtimetrains)",
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// doJSON performs a GET and decodes JSON into v.
func (c *Client) doJSON(ctx context.Context, pth string, v any) error {
	if ctx == nil {
		ctx = context.Background()
	}
	u := *c.baseURL
	// API paths already contain leading /json/...; if not, ensure we add /json
	// Accept caller passing either "service/..." or "/json/service/..." etc.
	clean := strings.TrimLeft(pth, "/")
	if !strings.HasPrefix(clean, "json/") {
		clean = path.Join("json", clean)
	}
	u.Path = path.Join(u.Path, clean)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.pass)
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return ErrUnauthorized
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return fmt.Errorf("rtt: unexpected status %d: %s", resp.StatusCode, strings.TrimSpace(string(b)))
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

// ErrNotFound returned when API 404s (e.g., invalid UID/date combination or station code).
var ErrNotFound = errors.New("rtt: not found")

// ErrUnauthorized returned on 401/403.
var ErrUnauthorized = errors.New("rtt: unauthorized")

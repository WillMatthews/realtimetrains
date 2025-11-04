package realtimetrains

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// SearchParams configures station search queries.
type SearchParams struct {
	To       string     // optional destination filter
	Date     *time.Time // optional date (YYYY/MM/DD)
	TimeHHMM string     // optional time in HHmm (only valid with Date)
	Arrivals bool       // if true, fetch arrivals rather than departures
}

// validate ensures params are coherent.
func (sp *SearchParams) validate() error { return nil } // currently permissive

// SearchStation retrieves departures (default) or arrivals for a station with optional filters.
func (c *Client) SearchStation(ctx context.Context, station string, p SearchParams) (*Container, error) {
	// build path segments
	segs := []string{"search", station}
	if p.To != "" {
		segs = append(segs, "to", p.To)
	}
	if p.Date != nil {
		y, m, d := p.Date.Date()
		segs = append(segs, fmt.Sprintf("%04d", y), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", d))
		if p.TimeHHMM != "" {
			segs = append(segs, p.TimeHHMM)
		}
	}
	if p.Arrivals {
		segs = append(segs, "arrivals")
	}
	path := strings.Join(segs, "/")
	var cont Container
	if err := c.doJSON(ctx, path, &cont); err != nil {
		return nil, err
	}
	return &cont, nil
}

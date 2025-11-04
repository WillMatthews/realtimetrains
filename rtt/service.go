package rtt

import (
    "context"
    "fmt"
    "time"
)

// GetService fetches a specific service for a given run date.
// date must be the service run date in local UK date (the API expects /YYYY/MM/DD/).
func (c *Client) GetService(ctx context.Context, serviceUID string, date time.Time) (*Service, error) {
    y, m, d := date.Date()
    p := fmt.Sprintf("service/%s/%04d/%02d/%02d", serviceUID, y, int(m), d)
    var svc Service
    if err := c.doJSON(ctx, p, &svc); err != nil { return nil, err }
    return &svc, nil
}

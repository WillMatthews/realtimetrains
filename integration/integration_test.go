//go:build integration
// +build integration

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/WillMatthews/realtimetrains"
)

const (
	londonCRS    = "KGX"
	cambridgeCRS = "CBG"
)

// newClient constructs a client using credentials from the environment.
func newClient(t *testing.T) *realtimetrains.Client {
	t.Helper()

	user := os.Getenv("RTT_USER")
	pass := os.Getenv("RTT_PASSWORD")
	if user == "" || pass == "" {
		t.Skip("set RTT_USER and RTT_PASSWORD to run integration tests")
	}

	client, err := realtimetrains.New(user, pass)
	if err != nil {
		t.Fatalf("realtimetrains.New: %v", err)
	}

	return client
}

func TestLondonToCambridgeSearch(t *testing.T) {
	// t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := newClient(t)

	params := realtimetrains.SearchParams{To: cambridgeCRS}
	services, err := client.SearchStation(ctx, londonCRS, params)
	if err != nil {
		t.Fatalf("SearchStation: %v", err)
	}

	if len(services.Services) == 0 {
		t.Fatalf("expected services from %s to %s", londonCRS, cambridgeCRS)
	}

	t.Logf("found %d services from %s to %s", len(services.Services), londonCRS, cambridgeCRS)
}

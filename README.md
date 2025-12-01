# realtimetrains-go

[![Go Tests](https://github.com/WillMatthews/realtimetrains/actions/workflows/go.yml/badge.svg)](https://github.com/WillMatthews/realtimetrains/actions/workflows/go.yml)

Minimal Go client for the RealTimeTrains (RTT) Pull / "simple" JSON API.

Official Pull API documentation: https://www.realtimetrains.co.uk/about/developer/pull/docs/

## Install

```
go get github.com/WillMatthews/realtimetrains
```

## Usage

```go
import (
    "context"
    "log"
    "time"
    rtt "github.com/WillMatthews/realtimetrains"
)

func example() {
    client, err := rtt.New("username", "password")
    if err != nil { log.Fatal(err) }

    // Search departures from KGX today
    ctx := context.Background()
    cont, err := client.SearchStation(ctx, "KGX", rtt.SearchParams{})
    if err != nil { log.Fatal(err) }
    log.Printf("services: %d", len(cont.Services))

    // Fetch a specific service (uid + date)
    date := time.Now()
    svc, err := client.GetService(ctx, cont.Services[0].ServiceUID, date)
    if err != nil { log.Fatal(err) }
    log.Printf("service has %d locations", len(svc.Locations))
}
```

## Integration tests

Live integration tests are kept under `integration/` and are guarded by the `integration` build tag so they do not run with the default `go test ./...`. Provide RTT credentials via the `RTT_USER` and `RTT_PASSWORD` environment variables, then run:

```
RTT_USER=example RTT_PASSWORD=secret go test -tags integration ./integration
```

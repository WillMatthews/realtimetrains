# gorealtimetrains

Minimal Go client for the RealTimeTrains (RTT) Pull / "simple" JSON API.

Official Pull API documentation: https://www.realtimetrains.co.uk/about/developer/pull/docs/

## Install

```
go get github.com/WillMatthews/gorealtimetrains
```

## Usage

```go
import (
    "context"
    "log"
    "time"
    "github.com/WillMatthews/gorealtimetrains/rtt"
)

func example() {
    client, err := rtt.New("username", "password")
    if err != nil { log.Fatal(err) }

    // Search departures from WAT today
    ctx := context.Background()
    cont, err := client.SearchStation(ctx, "WAT", rtt.SearchParams{})
    if err != nil { log.Fatal(err) }
    log.Printf("services: %d", len(cont.Services))

    // Fetch a specific service (uid + date)
    date := time.Now()
    svc, err := client.GetService(ctx, cont.Services[0].ServiceUID, date)
    if err != nil { log.Fatal(err) }
    log.Printf("service has %d locations", len(svc.Locations))
}
```

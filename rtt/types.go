package rtt

// NOTE: These structs model only the documented fields for the simple pull API.
// Many fields are left as basic types (string/bool/int) for now; helpers can be
// added later for richer representations.

// Pair represents an origin/destination pair object.
type Pair struct {
    TIPLOC       string `json:"tiploc"`
    Description  string `json:"description"`
    WorkingTime  string `json:"workingTime"`
    PublicTime   string `json:"publicTime"`
}

// Location represents a calling (or pass) point in a service or search result.
type Location struct {
    RealtimeActivated             bool    `json:"realtimeActivated"`
    TIPLOC                        string  `json:"tiploc"`
    CRS                           *string `json:"crs"`
    Description                   string  `json:"description"`
    WTTBookedArrival              *string `json:"wttBookedArrival"`
    WTTBookedDeparture            *string `json:"wttBookedDeparture"`
    WTTBookedPass                 *string `json:"wttBookedPass"`
    GBTTBookedArrival             *string `json:"gbttBookedArrival"`
    GBTTBookedDeparture           *string `json:"gbttBookedDeparture"`
    Origin                        []Pair  `json:"origin"`
    Destination                   []Pair  `json:"destination"`
    IsCall                        bool    `json:"isCall"`
    IsCallPublic                  bool    `json:"isCallPublic"`
    RealtimeArrival               *string `json:"realtimeArrival"`
    RealtimeArrivalActual         *bool   `json:"realtimeArrivalActual"`
    RealtimeArrivalNoReport       *bool   `json:"realtimeArrivalNoReport"`
    RealtimeWTTArrivalLateness    *int    `json:"realtimeWttArrivalLateness"`
    RealtimeGBTTAarrivalLateness  *int    `json:"realtimeGbttArrivalLateness"`
    RealtimeDeparture             *string `json:"realtimeDeparture"`
    RealtimeDepartureActual       *bool   `json:"realtimeDepartureActual"`
    RealtimeDepartureNoReport     *bool   `json:"realtimeDepartureNoReport"`
    RealtimeWTTDepartureLateness  *int    `json:"realtimeWttDepartureLateness"`
    RealtimeGBTTDepartureLateness *int    `json:"realtimeGbttDepartureLateness"`
    RealtimePass                  *string `json:"realtimePass"`
    RealtimePassActual            *bool   `json:"realtimePassActual"`
    RealtimePassNoReport          *bool   `json:"realtimePassNoReport"`
    Platform                      *string `json:"platform"`
    PlatformConfirmed             *bool   `json:"platformConfirmed"`
    PlatformChanged               *bool   `json:"platformChanged"`
    Line                          *string `json:"line"`
    LineConfirmed                 *bool   `json:"lineConfirmed"`
    Path                          *string `json:"path"`
    PathConfirmed                 *bool   `json:"pathConfirmed"`
    CancelReasonCode              *string `json:"cancelReasonCode"`
    CancelReasonShortText         *string `json:"cancelReasonShortText"`
    CancelReasonLongText          *string `json:"cancelReasonLongText"`
    DisplayAs                     *string `json:"displayAs"`
    ServiceLocation               *string `json:"serviceLocation"`
}

// Service returned by /service endpoint.
type Service struct {
    ServiceUID          string     `json:"serviceUid"`
    RunDate             string     `json:"runDate"`
    ServiceType         string     `json:"serviceType"`
    IsPassenger         bool       `json:"isPassenger"`
    TrainIdentity       string     `json:"trainIdentity"`
    PowerType           *string    `json:"powerType"`
    TrainClass          *string    `json:"trainClass"`
    Sleeper             *string    `json:"sleeper"`
    ATOCCode            string     `json:"atocCode"`
    ATOCName            string     `json:"atocName"`
    PerformanceMonitored bool      `json:"performanceMonitored"`
    Origin              []Pair     `json:"origin"`
    Destination         []Pair     `json:"destination"`
    Locations           []Location `json:"locations"`
    RealtimeActivated   bool       `json:"realtimeActivated"`
    RunningIdentity     *string    `json:"runningIdentity"`
}

// LocationContainer associates a location detail with service metadata (in search results).
type LocationContainer struct {
    LocationDetail  Location `json:"locationDetail"`
    ServiceUID      string   `json:"serviceUid"`
    RunDate         string   `json:"runDate"`
    TrainIdentity   string   `json:"trainIdentity"`
    RunningIdentity string   `json:"runningIdentity"`
    ATOCCode        string   `json:"atocCode"`
    ATOCName        string   `json:"atocName"`
    ServiceType     string   `json:"serviceType"`
    IsPassenger     bool     `json:"isPassenger"`
    PlannedCancel   bool     `json:"plannedCancel"`
    Origin          []Pair   `json:"origin"`
    Destination     []Pair   `json:"destination"`
    CountdownMinutes *int    `json:"countdownMinutes"`
}

// Container is the top-level response for search endpoints.
type Container struct {
    Location Location        `json:"location"`
    Filter   *SearchFilter   `json:"filter"`
    Services []LocationContainer `json:"services"`
}

// SearchFilter holds from/to filter data when supplied.
type SearchFilter struct {
    From *Location `json:"from"`
    To   *Location `json:"to"`
}

package realtimetrains

import (
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshalSearchSample(t *testing.T) {
	b, err := os.ReadFile("testdata/search_kgx_cbg.json")
	if err != nil {
		t.Fatal(err)
	}
	var c Container
	if err := json.Unmarshal(b, &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if *c.Location.CRS != "KGX" {
		t.Fatalf("location CRS mismatch: %+v", c.Location)
	}
	if len(c.Services) == 0 {
		t.Fatal("expected services")
	}
	first := c.Services[0]
	if first.LocationDetail.GBTTBookedDeparture == nil || first.LocationDetail.GBTTBookedDeparture.String() != "2027" {
		t.Fatalf("unexpected departure time: %#v", first.LocationDetail.GBTTBookedDeparture)
	}
	if len(first.LocationDetail.Associations) == 0 {
		t.Fatalf("expected association")
	}
}

func TestUnmarshalServiceSample(t *testing.T) {
	b, err := os.ReadFile("testdata/service_Y31309.json")
	if err != nil {
		t.Fatal(err)
	}
	var svc Service
	if err := json.Unmarshal(b, &svc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if svc.ServiceUID != "Y31309" {
		t.Fatalf("svc uid mismatch: %s", svc.ServiceUID)
	}
	if len(svc.Locations) < 2 {
		t.Fatalf("expected >=2 locations")
	}
	if svc.Locations[0].GBTTBookedDeparture == nil || svc.Locations[0].GBTTBookedDeparture.String() != "1926" {
		t.Fatalf("departure mismatch")
	}
}

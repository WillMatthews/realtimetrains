package realtimetrains

import (
	"encoding/json"
	"testing"
)

func TestTimeHHMM_Unmarshal(t *testing.T) {
	var v struct {
		T *TimeHHMM `json:"t"`
	}
	if err := json.Unmarshal([]byte(`{"t":"0935"}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.T == nil || v.T.Hour() != 9 || v.T.Minute() != 35 {
		t.Fatalf("unexpected %#v", v.T)
	}
	if v.T.String() != "0935" {
		t.Fatalf("string mismatch %s", v.T.String())
	}
}

func TestTimeHHMMSS_Unmarshal(t *testing.T) {
	var v struct {
		T *TimeHHMMSS `json:"t"`
	}
	if err := json.Unmarshal([]byte(`{"t":"235959"}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.T == nil || v.T.Hour() != 23 || v.T.Minute() != 59 || v.T.Second() != 59 {
		t.Fatalf("unexpected %#v", v.T)
	}
	if v.T.String() != "235959" {
		t.Fatalf("string mismatch %s", v.T.String())
	}
}

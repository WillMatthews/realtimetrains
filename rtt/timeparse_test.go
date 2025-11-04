package rtt

import "testing"

func TestParseHHMM(t *testing.T) {
    h,m,err := ParseHHMM("2359")
    if err != nil || h!=23 || m!=59 { t.Fatalf("unexpected %d %d %v", h,m,err) }
    if _,_,e := ParseHHMM("2400"); e==nil { t.Fatal("expected error") }
}

func TestParseHHMMSS(t *testing.T) {
    h,m,s,err := ParseHHMMSS("000000")
    if err!=nil || h!=0 || m!=0 || s!=0 { t.Fatalf("unexpected %d %d %d %v", h,m,s,err) }
    if _,_,_,e := ParseHHMMSS("999999"); e==nil { t.Fatal("expected error") }
}

package gojoursferies_test

import ("testing"
		"gojoursferies")

func TestZones(t *testing.T) {
	if len(Zones()) != 13 {
		t.Errorf("Zones() should return a slice of str with a length of 13")
	}
}

func TestCheckZone(t *testing.T) {
	z := Zones()[1]
	if res, _ :=CheckZone(z); res!= z && err==nil {
		t.Errorf("CheckZone(z) should return z on a valid zone")
	}
	if  resp, err :=CheckZone("This is an invalid string"); res== "" && err!=nil {
		t.Errorf("CheckZone should return an error on invalid zone")
	}
}

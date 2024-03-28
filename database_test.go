package main

import "testing"

func TestBuildDatabase(t *testing.T) {
	db, err := buildDatabase("testdata")
	if err != nil {
		t.Fatalf("buildDatabase() error = %v", err)
	}
	want := Profile{Path: "testdata/johndoe.ovpn", Password: "secret"}
	got, ok := db["johndoe"]
	if !ok {
		t.Fatalf("buildDatabase does not include expected profile: %#v", db)
	}
	if got != want {
		t.Fatalf("Database:\ngot:  %#v\nwant: %#v\n", got, want)
	}
}

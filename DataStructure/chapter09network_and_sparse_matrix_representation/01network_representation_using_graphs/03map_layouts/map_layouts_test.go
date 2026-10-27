package main

import "testing"

func TestNewMapLayout(t *testing.T) {
	mapLayout := NewMapLayout()
	t.Log(mapLayout)

	if mapLayout == nil {
		t.Errorf("error in creating a mapLayout")
	}
}

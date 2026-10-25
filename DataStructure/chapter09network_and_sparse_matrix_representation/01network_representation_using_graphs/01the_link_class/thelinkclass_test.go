package main

import "testing"

func TestNewSocialGraph(t *testing.T) {
	var socialGraph *SocialGraph
	socialGraph = NewSocialGraph(1)

	if socialGraph == nil {
		t.Errorf("error in creating a social Graph")
	}
}

package main

import "testing"

func TestNewKnowledgeGraph(t *testing.T) {
	knowledgeGraph := NewKnowledgeGraph()
	t.Log(knowledgeGraph)

	if knowledgeGraph == nil {
		t.Errorf("error in creating a knowledgeGraph")
	}

}

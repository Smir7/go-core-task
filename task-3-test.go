package main

import (
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	m.Add("rose", 1)
	if _, exists := m.Get("rose"); !exists {
		t.Errorf("Expected key 'rose' to exist with value 1, got %d")
	}

	if value, exists := m.Exist("rose"); !exists {
		t.Errorf("Expected key 'rose' to exist with value 1, got %d", value)
	}

	if value, exists := m.Get("rain"); exists {
		t.Errorf("Expected key 'rain' to not exist, got value %d", value)
	}

	m.Remove("rose")
	if _, exists := m.Get("rose"); exists {
		t.Error("Expected key 'rose' to be removed")
	}

	m.Add("tulip", 2)
	copyMap := m.Copy()
	if value, exists := copyMap.Exist("tulip"); !exists || value != 2 {
		t.Errorf("Expected key 'tulip' in copied map with value 2, got %d", value)
	}

	m.Remove("tulip")
	if _, exists := copyMap.Get("tulip"); !exists {
		t.Error("Expected key 'tulip' to exist in copied map")
	}
}

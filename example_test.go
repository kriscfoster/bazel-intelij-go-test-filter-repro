package main

import (
	"log"
	"testing"
)

func TestMultipleCases(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "TestCase1",
		},
		{
			name: "TestCase2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			log.Println("running test: " + tc.name)
		})
	}
}

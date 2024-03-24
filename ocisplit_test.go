package main

import (
	"log"
	"testing"
)

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

// Almost the same as the above, but this one is for single test instead of collection of tests
func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}

func TestMain(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	table := []struct {
		name     string
		expected bool
	}{
		{"one", true},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			actual := OCISplit(".", ".")
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}

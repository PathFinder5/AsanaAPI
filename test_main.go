// main_test.go
package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fetchDataCallCount int
func TestfetchData(){
	fetchData()
	fetchDataCallCount ++
}
// Test StartFetching function with assertions
func TestStartFetching(t *testing.T) {
	// Reset the call count before the test
	fetchDataCallCount := 0

	// Test for 1 minutes
	timer := 1
	go StartFetching(timer)

	// Run for a short time to allow the ticker to trigger a few times
	time.Sleep(2 * time.Minute) // Sleep longer than 10 seconds for a few ticks

	// Assert that fetchData was called at least once
	assert.Greater(t, fetchDataCallCount, 0, "fetchData should be called at least once")
}

// // Test for invalid timer input
// func TestInvalidTimer(t *testing.T) {
// 	// Normally this would need to handle invalid input gracefully,
// 	// since the original StartFetching function doesn't handle that,
// 	// let's create a small function for that.
// 	result := validateTimer(10) // Invalid timer

// 	// Assert that the result indicates an invalid timer
// 	assert.False(t, result, "Expected false for invalid timer")
// }

// // validateTimer checks if the timer is valid (5 or 30)
// func validateTimer(timer int) bool {
// 	return timer == 5 || timer == 30
// }

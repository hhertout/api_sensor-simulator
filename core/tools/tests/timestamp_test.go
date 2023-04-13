package tools_test

import (
	"api_sensor/core/tools"
	"strconv"
	"testing"
	"time"
)

func TestGetTimestamp(t *testing.T) {
	// Test that the timestamp is greater than 0
	t.Run("Positive", func(t *testing.T) {
		ts := tools.GetTimestamp()
		if ts <= 0 {
			t.Errorf("GetTimestamp() = %d; want > 0", ts)
		}
	})

	// Test that the timestamp is within a reasonable range
	t.Run("WithinRange", func(t *testing.T) {
		ts := tools.GetTimestamp()
		now := time.Now().UnixNano()
		if ts < now-1000000000 || ts > now+1000000000 {
			t.Errorf("GetTimestamp() = %d; want within %d of %d", ts, 1000000000, now)
		}
	})
}

func TestGetStringTimestamp(t *testing.T) {
	// Test that the string timestamp is not empty
	t.Run("NotEmpty", func(t *testing.T) {
		ts := tools.GetStringTimestamp()
		if ts == "" {
			t.Error("GetStringTimestamp() returned an empty string; want non-empty string")
		}
	})

	// Test that the string timestamp can be converted back to an int64
	t.Run("Parseable", func(t *testing.T) {
		ts := tools.GetStringTimestamp()
		i, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			t.Errorf("Error parsing timestamp string %q: %v", ts, err)
		}
		if i <= 0 {
			t.Errorf("Invalid timestamp %d parsed from string %q; want > 0", i, ts)
		}
	})

	// Test that the string timestamp corresponds to the current time within a reasonable range
	t.Run("CurrentTime", func(t *testing.T) {
		ts := tools.GetStringTimestamp()
		now := time.Now().UnixNano()
		i, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			t.Errorf("Error parsing timestamp string %q: %v", ts, err)
		}
		if i < now-1000000000 || i > now+1000000000 {
			t.Errorf("Invalid timestamp %d parsed from string %q; want within %d of %d", i, ts, 1000000000, now)
		}
	})
}

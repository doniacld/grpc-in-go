package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {

	now := time.Now()

	tests := []struct {
		name    string
		date    time.Time
		mood    Mood
		want    LogEntry
		wantErr error
	}{
		{name: "nominal", date: now, mood: Mood{state: "happy"}, want: LogEntry{date: now, mood: Mood{"happy"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewLogEntry(tt.date, tt.mood)
			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {

	testDate := time.Date(1993, time.November, 10, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name    string
		date    time.Time
		mood    Mood
		want    LogEntry
		wantErr error
	}{
		{name: "nominal", date: testDate, mood: Mood{State: "happy"}, want: LogEntry{date: testDate, mood: Mood{"happy"}}},
		{name: "error case", date: time.Now().Add(time.Hour * 12), mood: Mood{State: "happy"}, wantErr: ErrInvalidDate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLogEntry(tt.date, tt.mood)
			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

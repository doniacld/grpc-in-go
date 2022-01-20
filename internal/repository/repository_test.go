package repository

import (
	"context"
	"github.com/doniacld/grpc-in-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_SaveLog(t *testing.T) {

	logEntry, err := domain.NewLogEntry(time.Now(), domain.Mood{State: "happy"})
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		name    string
		entry   domain.LogEntry
		wantErr error
	}{
		{name: "nominal case", entry: logEntry},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			err := r.SaveLog(context.Background(), tt.entry)
			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	// TODO implement me
	t.Fail()
}

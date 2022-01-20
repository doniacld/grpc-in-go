package repository

import (
	"context"
	"fmt"
	"github.com/doniacld/grpc-in-go/internal/domain"
)

// Repository connects to the db to manage log entries.
type Repository struct{}

// SaveLog implements domain.LogRepository.
func (r Repository) SaveLog(ctx context.Context, entry domain.LogEntry) error {
	fmt.Printf("saving the log entry %v", entry)

	return nil
}

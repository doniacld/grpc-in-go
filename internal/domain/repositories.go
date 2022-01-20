package domain

import "context"

// LogRepository for logs.
type LogRepository interface {
	// SaveLog saves a log.
	SaveLog(context.Context, LogEntry) error
}

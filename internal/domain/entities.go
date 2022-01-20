package domain

import "time"

// ErrInvalidDate is returned when the date is in the future.
const ErrInvalidDate = Error("date is in the future")

// LogEntry is a daily entry.
type LogEntry struct {
	date time.Time
	mood Mood
}

// NewLogEntry creates a new log entry.
func NewLogEntry(date time.Time, mood Mood) (LogEntry, error) {
	if date.After(time.Now()) {
		return LogEntry{}, ErrInvalidDate
	}

	return LogEntry{
		date: date,
		mood: mood,
	}, nil
}

// MoodState is a state of mind.
type MoodState string

// Mood defines the mood log.
type Mood struct {
	state MoodState
}

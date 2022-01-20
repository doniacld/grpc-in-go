package domain

// Error implements error.
type Error string

// Error implements error.
func (e Error) Error() string {
	return string(e)
}

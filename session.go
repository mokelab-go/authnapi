package authnapi

// Session has Account ID
type Session interface {
	// ID is session ID
	ID() string
	// AccountID
	AccountID() string
}

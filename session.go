package authnapi

// Session has Account ID
type Session interface {
	// AccountID
	AccountID() string
}

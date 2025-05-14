package auth

// Update tokens
func Refresh() (*Tokens, error) {
	return &Tokens{
		Access:  "access",
		Refresh: "refresh",
	}, nil
}

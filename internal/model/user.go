package model

// UserAccount is an account with username and password.
//
// It is used to represent a user account in the system.
// The username is a unique identifier for the user, and the password is used for authentication.
//
// This struct is used for demonstration purposes only.
// In production, the password should be hashed and salted for security.
type UserAccount struct {
	Username string
	Password string
}

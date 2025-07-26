package abstraction

// User represents a user entity.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	// GetUserByID retrieves a user by their unique identifier.
	GetUserByID(id int) (*User, error)

	// CreateUser creates a new user record.
	CreateUser(user *User) error
}

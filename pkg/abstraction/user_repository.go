package abstraction

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

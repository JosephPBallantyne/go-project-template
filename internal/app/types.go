package app

type User struct {
	ID      int
	Name    string
	Address string
}

type UserService interface {
	// GetUser(id int) (*User, error)
	// GetUsers() ([]*User, error)
	CreateUser(u *User) error
	// DeleteUser(id int) error
}

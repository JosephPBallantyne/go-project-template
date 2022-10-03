package app

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UserService interface {
	GetUser(id int) (map[string]interface{}, error)
	// GetUser(id int) (*User, error)
	// GetUsers() ([]*User, error)
	CreateUser(u *User) error
	// DeleteUser(id int) error
}

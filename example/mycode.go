package example

import "fmt"

// MyService is an example mock service to work as demonstration support for
// Biff library.
type MyService struct {
	users    map[string]*User
	comments []string

	logged *User
}

// User represents a mock user for `MyService`
type User struct {

	// Email is the user email address without restrictions
	Email string

	// Password is the user clear password
	Password string
}

// RegisterUser will register a new user with a password. If user already
// exists a nil value will be returned.
func (m *MyService) RegisterUser(email, password string) *User {

	if _, exists := m.users[email]; exists {
		// User already exists
		return nil
	}

	u := &User{
		Email:    email,
		Password: password,
	}

	m.users[email] = u

	return u
}

// RetrieveUser will find a user by email. If user do not exist, nil will be
// returned.
func (m *MyService) RetrieveUser(email string) *User {
	if u, exists := m.users[email]; exists {
		return u
	}

	return nil
}

// Login will find a user by email, check password and put in session. If all
// the process is right, logged user will be returned, otherwise nil.
func (m *MyService) Login(email, password string) *User {

	u := m.RetrieveUser(email)
	if u == nil {
		return nil
	}

	if u.Password != password {
		return nil
	}

	m.logged = u

	return u
}

// WriteComment add a comment to the system. User should be logged in.
func (m *MyService) WriteComment(comment string) *string {

	if m.logged == nil {
		return nil
	}

	c := fmt.Sprintf("<%s> says: %s", m.logged.Email, comment)
	m.comments = append(m.comments, c)

	return &c
}

// NewMyService will instantiate a new `MyService`
func NewMyService() *MyService {
	return &MyService{
		users:    map[string]*User{},
		comments: []string{},
	}
}

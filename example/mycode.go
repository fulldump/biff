package example

import "fmt"

type MyService struct {
	users    map[string]*User
	comments []string

	logged *User
}

type User struct {
	Email    string
	Password string
}

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

func (m *MyService) RetrieveUser(email string) *User {
	if u, exists := m.users[email]; exists {
		return u
	}

	return nil
}

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

func (m *MyService) WriteComment(comment string) *string {

	if m.logged == nil {
		return nil
	}

	c := fmt.Sprintf("<%s> says: %s", m.logged.Email, comment)
	m.comments = append(m.comments, c)

	return &c
}

func NewMyService() *MyService {
	return &MyService{
		users:    map[string]*User{},
		comments: []string{},
	}
}

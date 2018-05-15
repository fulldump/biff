package example

import (
	"biff"
	"testing"
)

func TestMyCode(t *testing.T) {

	biff.Alternative("Instance service", func(a *biff.A) {

		s := NewMyService()
		a.AssertNotNil(s)

		a.Alternative("Retrieve unregistered user", func(a *biff.A) {
			user := s.RetrieveUser("john@email.com")
			a.AssertNil(user)
		})

		a.Alternative("Register user", func(a *biff.A) {

			john := s.RegisterUser("john@email.com", "john-123")
			a.AssertNotNil(john)

			a.Alternative("Register the same user twice ", func(a *biff.A) {
				john := s.RegisterUser("john@email.com", "john-123")
				a.AssertNil(john)
			})

			a.Alternative("Retrieve user", func(a *biff.A) {
				user := s.RetrieveUser("john@email.com")
				a.AssertNotNil(user)
				a.AssertEqual(user, &User{
					Email:    "john@email.com",
					Password: "john-123",
				})
			})

			a.Alternative("Bad credentials", func(a *biff.A) {
				user := s.Login("john@email.com", "bad-password")
				a.AssertNil(user)
			})

			a.Alternative("Write comment", func(a *biff.A) {
				comment := s.WriteComment("Hello world")
				a.AssertNil(comment)
			})

			a.Alternative("Login", func(a *biff.A) {
				a.SetDescription(`When a user is logged in, it should be
				returned back.`)

				user := s.Login("john@email.com", "john-123")
				a.AssertNotNil(user)

				a.Alternative("Write comment", func(a *biff.A) {
					comment := s.WriteComment("Hello world!")
					a.AssertNotNil(comment)
					a.AssertEqual(s.comments[0], "<john@email.com> says: Hello world!")
				})
			})
		})
	})

}

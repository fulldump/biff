package example

import (
	"biff"
	"testing"
)

func TestMyCode(t *testing.T) {

	biff.Alternative("Create user", func(a *biff.A) {

		a.Alternative("Retrieve user", func(a *biff.A) {

		})

		a.Alternative("Delete user", func(a *biff.A) {

		})

		a.Alternative("Ban user", func(a *biff.A) {

			a.Alternative("Unban user", func(a *biff.A) {

			})

			a.Alternative("Forbidden comment", func(a *biff.A) {

			})

		})

	})

}

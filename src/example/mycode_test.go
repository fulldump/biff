package example

import (
	"biff"
	"fmt"
	"testing"
)

func TestMyCode(t *testing.T) {

	biff.Alternative(func(a *biff.A) {
		fmt.Println("Create user")

		a.Alternative(func(a *biff.A) {
			fmt.Println("Print user")
		})

		a.Alternative(func(a *biff.A) {
			fmt.Println("Delete user")
		})

		a.Alternative(func(a *biff.A) {
			fmt.Println("Ban user")

			a.Alternative(func(a *biff.A) {
				fmt.Println("Unban user")
			})

			a.Alternative(func(a *biff.A) {
				fmt.Println("Forbidden comment")
			})

		})

	})

}

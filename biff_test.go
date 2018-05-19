package biff

import (
	"testing"
)

func Test_isolation(t *testing.T) {

	Alternative("test a", func(a *A) {

		words := []string{}

		words = append(words, a.Title)
		a.AssertEqual(words, []string{"test a"})

		a.Alternative("test a1", func(a *A) {
			words = append(words, a.Title)
			a.AssertEqual(words, []string{"test a", "test a1"})
		})

		a.Alternative("test a2", func(a *A) {
			words = append(words, a.Title)
			a.AssertEqual(words, []string{"test a", "test a2"})
		})

	})

}

func Test_scope(t *testing.T) {

	sequence := []string{"zero"}

	Alternative("test a", func(a *A) {

		sequence = append(sequence, "one")

		a.Alternative("test b", func(a *A) {

			sequence = append(sequence, "two")
			a.AssertEqual(sequence, []string{"zero", "one", "two"})
		})

		sequence := []string{"three"}

		a.Alternative("test c", func(a *A) {

			sequence = append(sequence, "four")
			a.AssertEqual(sequence, []string{"three", "four"})
		})

	})

}

func Example_basicUsage() {

	Alternative("Initial value", func(a *A) {
		value := 10
		a.AssertEqual(value, 10)

		a.Alternative("Plus 50", func(a *A) {
			// Here value == 10
			value += 50
			a.AssertEqual(value, 60)
		})

		a.Alternative("Multiply by 2", func(a *A) {
			// Here value == 10 again (it is an alternative from the parent)
			value *= 2
			a.AssertEqual(value, 20)
		})
	})

	// Output:
	// Case: Initial value
	//     value is 10
	// Case: Plus 50
	//     value is 60
	// -------------------------------
	// Case: Initial value
	//     value is 10
	// Case: Multiply by 2
	//     value is 20
	// -------------------------------
	// Case: Initial value
	//     value is 10
	// -------------------------------

}

func Test_trimMultiline(t *testing.T) {

	Alternative("TrimMultiline", func(a *A) {

		a.Alternative("empty", func(a *A) {

			a.AssertEqual("", trimMultiline(""))

		})

		a.Alternative("multiple lines", func(a *A) {

			a.AssertEqual(trimMultiline(`
one
    two
	     three       
`), "\none\ntwo\nthree\n\n")

		})

	})

}

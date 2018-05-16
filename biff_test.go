package biff

import "testing"

func Test_isolation(t *testing.T) {

	Alternative("test a", func(a *A) {

		words := []string{}

		words = append(words, a.title)
		a.AssertEqual(words, []string{"test a"})

		a.Alternative("test a1", func(a *A) {
			words = append(words, a.title)
			a.AssertEqual(words, []string{"test a", "test a1"})
		})

		a.Alternative("test a2", func(a *A) {
			words = append(words, a.title)
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

package biff

import "fmt"

type A struct {
	skip      int
	f         F
	substatus *[]int
	done      bool
}

func NewTest(f F) *A {
	return &A{
		f: f,
	}
}

func (t *A) Alternative(f F) *A {

	if t.skip == 0 {
		t.done = NewTest(f).Run(t.substatus)
	}

	t.skip--

	return t
}

func (t *A) Run(status *[]int) (done bool) {

	skip := &(*status)[0]
	substatus := (*status)[1:]

	// Execute
	t.skip = *skip
	t.substatus = &substatus
	t.f(t)

	// There is no more alternatives
	if t.skip == 0 {
		(*status)[0] = 0
		return true
	}

	if t.done {
		(*status)[0]++
	}

	return
}

type F func(t *A)

func Alternative(f F) {

	// Ñap :_(
	status := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for {

		t := NewTest(f)

		done := t.Run(&status)

		fmt.Println("-------------------------------")

		if done {
			return
		}

	}

}

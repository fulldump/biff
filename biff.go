package biff

import "fmt"

type A struct {
	skip      int
	f         F
	substatus *[]int
	done      bool
	title     string
}

func NewTest(f F) *A {
	return &A{
		f: f,
	}
}

func (t *A) Alternative(title string, f F) *A {

	if t.skip == 0 {
		n := NewTest(f)
		n.title = title
		t.done = n.Run(t.substatus)
	}

	t.skip--

	return t
}

func (t *A) Run(status *[]int) (done bool) {

	fmt.Println("Case:", t.title)

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

func Alternative(title string, f F) {

	// Ñap :_(
	status := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for {

		t := NewTest(f)
		t.title = title

		done := t.Run(&status)

		fmt.Println("-------------------------------")

		if done {
			return
		}

	}

}

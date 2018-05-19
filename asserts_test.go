package biff

func Example_JSON_Equality() {

	Alternative("Json equality", func(a *A) {

		i := map[string]interface{}{
			"number": int(33),
		}
		f := map[string]interface{}{
			"number": float64(33),
		}

		a.AssertEqualJson(i, f)

	})

	// Output:
	// Case: Json equality
	//     i is map[string]interface {}{"number":33}
	// -------------------------------
}

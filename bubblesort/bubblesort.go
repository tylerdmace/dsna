package bubblesort

func Bubblesort(l []int) []int {
	f := false

	for i := 0; i < len(l)-1; i++ {
		a := l[i]
		b := l[i+1]

		if a > b {
			l[i] = b
			l[i+1] = a
			f = true
		}
	}

	if !f {
		return l
	}

	return Bubblesort(l)
}

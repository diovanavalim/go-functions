package function

func MultipleOp(n1, n2 int8) (int8, int8) {
	var sum int8 = Sum(n1, n2)
	var sub int8 = Sub(n1, n2)

	return sum, sub
}

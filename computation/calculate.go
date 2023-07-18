package computation

func Calculate(a, b int, operator string) (res int) {
	switch operator {
	case "+":
		{
			res = a + b
		}
	case "-":
		{
			res = a - b
		}
	case "*":
		{
			res = a * b
		}
	case "/":
		{
			res = a / b
		}
	}

	return res
}

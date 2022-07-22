package challenge06

func ValidateParenthesesOrderOpt(k string) bool {
	counter := 0

	for _, char := range k {
		if char == '(' {
			counter++
		} else {
			counter--
		}

		if counter < 0 {
			return false
		}
	}

	return true
}

func ValidateParenthesesOrder(k string) bool {
	aux := []string{}

	for _, char := range k {
		if char == '(' {
			aux = append(aux, "(")
		} else {
			if len(aux) == 0 {
				return false
			}
			aux = aux[:len(aux)-1]
		}
	}

	return len(aux) == 0
}

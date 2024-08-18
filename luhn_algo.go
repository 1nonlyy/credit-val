package main

func luhnAlgo(num string) bool {
	total := 0
	isSecondDig := false


	for i:= len(num) - 1; i>=0; i-- {
		digit := int(num[i] - '0')
		if isSecondDig {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}

		}
		total += digit 
		isSecondDig = !isSecondDig
	}

	return total%10 == 0
}
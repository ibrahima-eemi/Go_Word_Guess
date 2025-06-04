package exo1

// 1. Somme de trois nombres
func Sum(a, b, c int) int {
	return a + b + c
}

// 2. Nombre pair
func IsEven(n int) bool {
	return n%2 == 0
}

// 3. Maximum de quatre nombres
func MaxOfFour(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

// 4. Calcul de la factorielle
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 5. Nombre d'occurrences d'un caractère
func CountOccurrences(s string, r rune) int {
	count := 0
	for _, char := range s {
		if char == r {
			count++
		}
	}
	return count
}

// 6. Filtrer les nombres pairs
func FilterEven(numbers []int) []int {
	evens := make([]int, 0)
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}


// 7. Inverser une chaîne de caractères
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

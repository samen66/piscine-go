package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	sum := 0
	Collatz(&start, &sum)

	return sum
}

func Collatz(n *int, sum *int) bool {
	if *n == 1 {
		return false
	}
	if *n%2 == 0 {
		*n = *n / 2
		*sum++
	} else {
		*n = *n*3 + 1
		*sum++
	}
	return Collatz(n, sum)
}

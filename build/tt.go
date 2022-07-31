package build

func Pow1(m, n int) int {
	if n == 1 {
		return m
	}
	tmp := Pow1(m, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return n % 2 * tmp * tmp
}

func Pow2(a, b int) int {
	r, base := 1, a
	for b != 0 {
		if b%2 != 0 {
			r *= base
		}
		base *= base
		b /= 2
	}
	return r
}

func Pow3(x, n int) int {
	if n == 0 {
		return 1
	}

	t := 1
	for n != 0 {
		if n&1 != 0 {
			t *= x
		}
		n >>= 1
		x *= x
	}
	return t
}

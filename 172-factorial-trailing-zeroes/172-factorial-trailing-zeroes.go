package leetcode

import "math/big"

// Problem: 172. Factorial Trailing Zeroes
// Link: https://leetcode.com/problems/factorial-trailing-zeroes/description/
// Lessons lerned:
// - Never work with math/big package => My first solution don't work with n = 30
// - Try math/big package and failed on n = 6079 cause time limit is reached
// - 
func trailingZeroes_with_big(n int) int {
	n_factorial := big.NewInt(1)
	for i := 2; i <= n; i++ {
		n_factorial.Mul(n_factorial, big.NewInt(int64(i)))
	}

	ten := big.NewInt(10)
	zero := big.NewInt(0)

	zeroes := 0
	mod_res := new(big.Int)
	for {
		mod_res.Mod(n_factorial, ten)
		if mod_res.Cmp(zero) == 0 {
			n_factorial.Div(n_factorial, ten)
			zeroes++
		} else {
			return zeroes
		}
	}
}

package leetcode

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return fastpow(x, n)
}

func fastpow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := fastpow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

// @lc code=end

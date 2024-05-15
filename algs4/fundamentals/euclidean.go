package fundamentals

// GCD (Greatest Common Divisor) 「欧几里得算法」求两个非负整数 p 和 q 的最大公约数：
//   - 若 q 是 0，则最大公约数为 p；
//   - 否则，将 p 除以 q 得到余数 r，p 和 q 的最大公约数即为 q 和 r 的最大公约数。
func GCD(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return GCD(q, r)
}

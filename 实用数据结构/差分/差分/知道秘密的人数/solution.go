package leetcode_2327

// https://leetcode.cn/problems/number-of-people-aware-of-a-secret/
// 难度：中等

const mod = 1000000007

func peopleAwareOfSecret(n int, delay int, forget int) int {
	known, share := make([]int, n+forget+1), make([]int, n+forget+1)
	known[1], known[1+forget] = 1, -1
	share[1+delay], share[1+forget] = 1, -1

	sum := 0
	for i := 1; i <= n; i++ {
		sum = (sum + share[i] + mod) % mod

		known[i] += sum
		known[i+forget] -= sum
		share[i+delay] += sum
		share[i+forget] -= sum
	}

	sum = 0
	for i := 1; i <= n; i++ {
		sum += known[i]
	}

	return sum % mod
}

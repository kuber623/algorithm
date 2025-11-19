package leetcode_2559

// https://leetcode.cn/problems/count-vowel-strings-in-ranges/
// 难度：中等
// 题解：前缀和

func vowelStrings(words []string, queries [][]int) (ans []int) {
	ans = make([]int, len(queries))

	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	// prefsum[i] 表示 words 数组中 [0, i-1] 区间中元音字符串数的累加和
	prefsum := make([]int, len(words)+1)
	for i, word := range words {
		added := 0
		if vowel[word[0]] && vowel[word[len(word)-1]] {
			added = 1
		}
		prefsum[i+1] = prefsum[i] + added
	}

	for i := 0; i < len(queries); i++ {
		ans[i] = prefsum[queries[i][1]] - prefsum[queries[i][0]]
	}

	return
}

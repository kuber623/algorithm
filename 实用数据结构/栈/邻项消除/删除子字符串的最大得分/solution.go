package leetcode_1717

// https://leetcode.cn/problems/maximum-score-from-removing-substrings/
// 难度：中等
// 题解：
//

func maximumGain(s string, x int, y int) (score int) {
	// 预处理字符串
	// 如果 ba 模式比 ab 模式价值更高则对字符和价值进行翻转
	bs := []byte(s)
	if x < y {
		// 翻转字符
		for i := 0; i < len(bs); i++ {
			if bs[i] == 'a' {
				bs[i] = 'b'
			} else if bs[i] == 'b' {
				bs[i] = 'a'
			}
		}
		// 翻转价值
		x, y = y, x
	}

	cnta, cntb := 0, 0 // 记录 ab 模式串中 a 和 b 的个数
	for i := 0; i < len(bs); i++ {
		switch bs[i] {
		case 'a':
			cnta++
		case 'b':
			if cnta > 0 {
				cnta--
				score += x
			} else {
				cntb++
			}
		default:
			score += min(cnta, cntb) * y
			cnta, cntb = 0, 0
		}
	}
	score += min(cnta, cntb) * y

	return
}

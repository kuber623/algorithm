package leetcode_3597

// https://leetcode.cn/problems/partition-string/
// 难度：中等

func partitionString(s string) (ans []string) {
	trie := &Trie{root: &Node{}}

	cur, segment := trie.root, make([]rune, 0)
	for _, ch := range s {
		segment = append(segment, ch)

		ch -= 'a'
		if cur.son[ch] != nil {
			cur = cur.son[ch]
		} else {
			cur.son[ch] = &Node{}
			ans = append(ans, string(segment))
			cur = trie.root
			segment = segment[:0]
		}
	}
	return
}

type Trie struct {
	root *Node
}

type Node struct {
	son [26]*Node
}

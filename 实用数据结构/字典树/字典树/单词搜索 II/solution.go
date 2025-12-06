package leetcode_212

func findWords(board [][]byte, words []string) (ans []string) {
	trie := &Trie{root: &Node{}}
	for _, word := range words {
		trie.Insert(word)
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			ch := board[i][j]
			visited[i][j] = true
			trackback(trie.root.son[ch-'a'], []byte{ch}, board, &visited, i, j, &ans)
			visited[i][j] = false
		}
	}

	// 单词去重
	m := make(map[string]bool, len(ans))
	for _, word := range ans {
		m[word] = true
	}
	ans = make([]string, 0, len(m))
	for word := range m {
		ans = append(ans, word)
	}

	return
}

func trackback(cur *Node, word []byte, board [][]byte, visited *[][]bool, i, j int, ans *[]string) {
	if cur == nil {
		return
	}

	m, n := len(board), len(board[0])

	if cur.end {
		*ans = append(*ans, string(word))
	}

	// 上
	if i > 0 && !(*visited)[i-1][j] {
		ch := board[i-1][j]
		if cur.son[ch-'a'] != nil {
			word = append(word, ch)
			(*visited)[i-1][j] = true
			trackback(cur.son[ch-'a'], word, board, visited, i-1, j, ans)
			(*visited)[i-1][j] = false
			word = word[:len(word)-1]
		}
	}
	// 下
	if i < m-1 && !(*visited)[i+1][j] {
		ch := board[i+1][j]
		if cur.son[ch-'a'] != nil {
			word = append(word, ch)
			(*visited)[i+1][j] = true
			trackback(cur.son[ch-'a'], word, board, visited, i+1, j, ans)
			(*visited)[i+1][j] = false
			word = word[:len(word)-1]
		}
	}
	// 左
	if j > 0 && !(*visited)[i][j-1] {
		ch := board[i][j-1]
		if cur.son[ch-'a'] != nil {
			word = append(word, ch)
			(*visited)[i][j-1] = true
			trackback(cur.son[ch-'a'], word, board, visited, i, j-1, ans)
			(*visited)[i][j-1] = false
			word = word[:len(word)-1]
		}
	}
	// 右
	if j < n-1 && !(*visited)[i][j+1] {
		ch := board[i][j+1]
		if cur.son[ch-'a'] != nil {
			word = append(word, ch)
			(*visited)[i][j+1] = true
			trackback(cur.son[ch-'a'], word, board, visited, i, j+1, ans)
			(*visited)[i][j+1] = false
			word = word[:len(word)-1]
		}
	}
}

type Trie struct {
	root *Node
}

type Node struct {
	son [26]*Node
	end bool
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			cur.son[ch] = &Node{}
		}
		cur = cur.son[ch]
	}
	cur.end = true
}

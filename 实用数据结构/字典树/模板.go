package trie

type Trie struct {
	root *Node
}

type Node struct {
	son [26]*Node
	end bool
}

func NewTrie() *Trie {
	return &Trie{&Node{}}
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

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			return false
		}
		cur = cur.son[ch]
	}
	return cur.end
}

func (t *Trie) StartWith(prefix string) bool {
	cur := t.root
	for _, ch := range prefix {
		ch -= 'a'
		if cur.son[ch] == nil {
			return false
		}
		cur = cur.son[ch]
	}
	return true
}

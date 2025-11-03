package aho_corasick

import "container/list"

type AhoCorasick struct {
	root *node
}

type node struct {
	children map[rune]*node // 子节点
	fail     *node          // 失败指针
	matches  []string       // 该节点对应模式串
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: &node{
			children: make(map[rune]*node),
		},
	}
}

func (ac *AhoCorasick) AddPattern(pattern string) {
	cur := ac.root
	for _, c := range pattern {
		if _, ok := cur.children[c]; !ok {
			cur.children[c] = &node{
				children: make(map[rune]*node),
			}
			cur = cur.children[c]
		}
	}
	cur.matches = append(cur.matches, pattern)
}

func (ac *AhoCorasick) BuildFailureLink() {
	queue := list.New()

	// 根节点的直接子节点（深度为 1 的节点）的失败指针指向根节点
	for _, child := range ac.root.children {
		child.fail = ac.root
		queue.PushBack(child)
	}

	// 广度优先遍历 Trie 树，构建失败指针，其中失败指针指向的「最长的、同时也是其他模式串前缀的后缀」的节点
	// 构建规则：
	// 假设父节点为 p，当前子节点为 c
	// 当节点 p 失败指针指向的节点 f 存在与节点 c 相同字符的子节点 v 时，c 的失败指针指向 v
	// 反之则更新 f = f.fail
	for queue.Len() > 0 {
		cur := queue.Front().Value.(*node)
		queue.Remove(queue.Front())

		for char, c := range cur.children {
			f := cur.fail
			for f != nil {
				if v, ok := f.children[char]; ok {
					c.fail = v
					break
				}
				f = f.fail
			}
			if f == nil {
				c.fail = ac.root
			}
			c.matches = append(c.matches, c.fail.matches...)

			queue.PushBack(c)
		}
	}
}

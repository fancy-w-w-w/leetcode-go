package tree

// 可以用map实现trieNode
type TrieV3 struct {
	children map[byte]*TrieV3
	word     string
}

func (t *TrieV3) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = &TrieV3{children: map[byte]*TrieV3{}}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// FindWords
// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
// dfs + 剪枝
func FindWords(board [][]byte, words []string) (ans []string) {
	t := &TrieV3{children: map[byte]*TrieV3{}}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])

	var dfs func(node *TrieV3, x, y int)
	dfs = func(node *TrieV3, x, y int) {
		ch := board[x][y]
		nxt := node.children[ch]
		if nxt == nil {
			return
		}

		// 去重
		if nxt.word != "" {
			ans = append(ans, nxt.word)
			nxt.word = ""
		}

		if len(nxt.children) > 0 {
			board[x][y] = '#'
			for _, d := range dirs {
				nx, ny := x+d.x, y+d.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
					dfs(nxt, nx, ny)
				}
			}
			board[x][y] = ch
		}

		// 剪枝
		if len(nxt.children) == 0 {
			delete(node.children, ch)
		}
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	return
}

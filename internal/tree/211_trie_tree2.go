package tree

// WordDictionary 字典树，实现能匹配通配符的"."的string
type WordDictionary struct {
	trieNode [26]*WordDictionary
	isEnd    bool
}

func NeWordDictionaryw() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.trieNode[ch] == nil {
			node.trieNode[ch] = &WordDictionary{}
		}
		node = node.trieNode[ch]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *WordDictionary) bool
	dfs = func(index int, node *WordDictionary) bool {
		if index == len(word) {
			return node.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.trieNode[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.trieNode {
				child := node.trieNode[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

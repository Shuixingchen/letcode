package algorithm

//前缀树，字典树

type trieNode struct {
	value     string             //表示一个utf8字符
	childrens map[rune]*trieNode //每个节点可能有多个子节点
	isEnd     bool
}
type trie struct {
	root *trieNode
}

func NewTrieNode(char string) *trieNode {
	return &trieNode{
		value:     char,
		childrens: make(map[rune]*trieNode),
		isEnd:     false,
	}
}

func NewTrie() *trie {
	return &trie{
		root: NewTrieNode("/"),
	}
}
func (t *trie) Insert(word string) {
	str := []rune(word)
	node := t.root
	for i := 0; i < len(str); i++ {
		code := str[i]
		childNode, ok := node.childrens[code]
		if !ok {
			node.childrens[code] = NewTrieNode(string(code))
		} else {
			node = childNode
		}
		node.isEnd = true
	}
}

//遍历trie树
func (t *trie) Traversal() {

}

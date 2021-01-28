package dataStruct
//前缀树，字典树

type trieNode struct {
	childrens map[interface{}]*trieNode
	isEnd bool
}
func NewTrieNode () *trieNode{
	return &trieNode{
		childrens: make(map[interface{}]*trieNode),
		isEnd:    false,
	}
}
type Trie struct {
	root *trieNode
}
func NewTrie () *Trie{
	return &Trie{root:NewTrieNode()}
}

func (t *Trie)Insert(word string) {
	node := t.root
	for i:=0;i<len(word);i++{
		_,ok := node.childrens[word[i]]
		if !ok {
			node.childrens[word[i]] = NewTrieNode()
		}
		node = node.childrens[word[i]]
	}
	node.isEnd = true
}

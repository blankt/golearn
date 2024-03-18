package main

/*
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/

func main() {
	root := Constructor()
	root.Insert("apple")
	root.Search("apple")
}

type Trie struct {
	isEnd bool
	Next  []*Trie
}

func Constructor() Trie {
	root := Trie{}
	root.Next = make([]*Trie, 26)
	return root
}

func (t *Trie) Insert(word string) {
	root := t
	for _, v := range word {
		index := v - 'a'
		if root.Next[index] == nil {
			next := &Trie{}
			next.Next = make([]*Trie, 26)
			root.Next[index] = next
			root = next
		} else {
			root = root.Next[index]
		}
	}
	root.isEnd = true
}

func (t *Trie) Search(word string) bool {
	root := t
	for _, v := range word {
		index := v - 'a'
		if root.Next[index] == nil {
			return false
		} else {
			root = root.Next[index]
		}
	}
	if root.isEnd {
		return true
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	root := t
	for _, v := range prefix {
		index := v - 'a'
		if root.Next[index] == nil {
			return false
		}
		root = root.Next[index]
	}
	return true
}

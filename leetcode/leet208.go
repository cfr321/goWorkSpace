package main

type Trie struct {
	son [26]*Trie
	end bool
}

/** Initialize your data structure here. */
func Constructor2() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	a := word[0]
	if this.son[a-'a'] == nil {
		trie := Constructor2()
		this.son[a-'a'] = &trie
	}
	if len(word) == 1 {
		this.son[a-'a'].end = true
	} else {
		this.son[a-'a'].Insert(word[1:])
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	a := word[0]
	if this.son[a-'a'] == nil {
		return false
	}
	if len(word) == 1 {
		return this.son[a-'a'].end
	} else {
		return this.son[a-'a'].Search(word[1:])
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	a := prefix[0]
	if this.son[a-'a'] == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	} else {
		return this.son[a-'a'].StartsWith(prefix[1:])
	}
}

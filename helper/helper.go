package helper

import "strings"

type Node struct {
	next          map[rune]*Node
	previous      *Node
	terminalCount int
	subtreeMax    int
	maxRune       rune
}

func NewNode(previous *Node) *Node {
	return &Node{make(map[rune]*Node), previous, 0, 0, ' '}
}

type Helper struct {
	rootNode *Node
	nowNode  *Node
	nowWord  string
}

func NewHelper() *Helper {
	n := NewNode(nil)
	h := &Helper{n, n, ""}
	return h
}

func (h *Helper) clear() {
	h.nowNode = h.rootNode
	h.nowWord = ""
}

func (h *Helper) AddString(s string) {
	h.clear()
	runes := []rune(s)

	n := h.rootNode
	for _, char := range runes {
		nn, ok := n.next[char]
		if !ok {
			n.next[char] = NewNode(n)
			nn = n.next[char]
		}

		n = nn
	}

	n.terminalCount++
	if n.subtreeMax < n.terminalCount {
		newMax := n.terminalCount
		n.subtreeMax = newMax

		for i := len(runes) - 1; i >= 0; i-- {
			n = n.previous

			if n.subtreeMax < newMax {
				n.subtreeMax = newMax
				n.maxRune = runes[i]
			} else {
				break
			}
		}
	}
}

func (h *Helper) AddText(text string) {
	for _, s := range strings.Fields(text) {
		h.AddString(s)
	}
}

func (h *Helper) MakeNewRequest(s string) (string, bool) {
	h.clear()
	runes := []rune(s)

	n := h.rootNode
	for _, char := range runes {
		nn, ok := n.next[char]
		if !ok {
			return "", false
		}

		n = nn
	}

	h.nowNode = n
	h.nowWord = s

	result := make([]rune, 0)
	for n.terminalCount < n.subtreeMax {
		result = append(result, n.maxRune)
		n, _ = n.next[n.maxRune]
	}

	return h.nowWord + string(result), true
}

func (h *Helper) ExpandRequest(s string) (string, bool) {
	runes := []rune(s)

	n := h.nowNode
	for _, char := range runes {
		nn, ok := n.next[char]
		if !ok {
			h.clear()
			return "", false
		}

		n = nn
	}

	h.nowNode = n
	h.nowWord += s

	result := make([]rune, 0)
	for n.terminalCount < n.subtreeMax {
		result = append(result, n.maxRune)
		n, _ = n.next[n.maxRune]
	}

	return h.nowWord + string(result), true
}

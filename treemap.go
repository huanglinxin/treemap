package treemap

import (
	"fmt"
	"strings"
)

type node struct {
	left   *node
	right  *node
	height int
	key    string
	value  string
}
type TreeMap struct {
	root *node
}

func maxHeight(left *node, right *node) int {
	leftH := 0
	rightH := 0

	if left != nil {
		leftH = left.height
	}
	if right != nil {
		rightH = right.height
	}
	if leftH > rightH {
		return leftH
	}
	return rightH
}
func treeBalance(p *node) bool {
	abs := p.left.height - p.right.height
	if abs < 0 {
		abs = -abs
	}
	if abs >= 2 {
		return false
	}
	return true
}
func countBalanceAlpha(p *node) int {
	if p == nil {
		return 0
	}
	leftH := 0
	rightH := 0
	if p.left != nil {
		leftH = p.left.height
	}
	if p.right != nil {
		rightH = p.right.height
	}
	return leftH - rightH
}
func rotateLeft(p *node) *node {
	rchild := p.right
	p.right = rchild.left
	rchild.left = p
	p.height = maxHeight(p.left, p.right) + 1
	rchild.height = maxHeight(rchild.left, rchild.right) + 1
	return rchild
}
func rotateRight(p *node) *node {
	lchild := p.left
	p.left = lchild.right
	lchild.right = p
	p.height = maxHeight(p.left, p.right) + 1
	lchild.height = maxHeight(lchild.left, lchild.right) + 1
	return lchild
}
func rotate(p *node) *node {
	if treeBalance(p) == true {
		return p
	}
	if countBalanceAlpha(p) == 2 {
		if countBalanceAlpha(p.left) == -1 {
			p.left = rotateLeft(p.left)
		}
		p = rotateRight(p)
		return p
	}
	if countBalanceAlpha(p) == -2 {
		if countBalanceAlpha(p.right) == 1 {
			p.right = rotateRight(p.right)
		}
		p = rotateLeft(p)
		return p
	}
	return p
}
func insertKeyAndValue(this *node, key string, value string) *node {
	if this == nil {
		return &node{left: nil, right: nil, height: 1, key: key, value: value}
	}
	if strings.Compare(key, this.key) == 0 {
		this.key = key
		return this
	}
	if strings.Compare(key, this.key) < 0 {
		this.left = insertKeyAndValue(this.left, key, value)
		this.height = maxHeight(this.left, this.right) + 1
		return rotate(this)
	}
	this.right = insertKeyAndValue(this.right, key, value)
	this.height = maxHeight(this.left, this.right) + 1
	return rotate(this)

}
func (this *TreeMap) Insert(key string, value string) {
	this.root = insertKeyAndValue(this.root, key, value)
	return
}
func (this *TreeMap) ShowAll() {
	LeftTree := TreeMap{root: this.root.left}
	rightTree := TreeMap{root: this.root.right}
	LeftTree.ShowAll()
	fmt.Println(this.root.key + " : " + this.root.value)
	rightTree.ShowAll()
}
func Helloworld(msg string) {
	fmt.Println("helloworld" + msg)
	return
}

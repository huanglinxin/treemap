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
	parent *node
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
	abs := countBalanceAlpha(p)
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
	//完成旋转
	rchild.parent = p.parent
	p.parent = rchild
	//更新parent指针
	if rchild.parent != nil && rchild.parent.left == p {
		rchild.parent.left = rchild
	}
	if rchild.parent != nil && rchild.parent.right == p {
		rchild.parent.right = rchild
	}
	//更新parent的child指针
	p.height = maxHeight(p.left, p.right) + 1
	rchild.height = maxHeight(rchild.left, rchild.right) + 1
	return rchild
}
func rotateRight(p *node) *node {
	lchild := p.left
	p.left = lchild.right
	lchild.right = p
	//完成旋转
	lchild.parent = p.parent
	p.parent = lchild
	//更新parent指针
	if lchild.parent != nil && lchild.parent.left == p {
		lchild.parent.left = lchild
	}
	if lchild.parent != nil && lchild.parent.right == p {
		lchild.parent.right = lchild
	}
	//更新parent的child指针
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
		this.left.parent = this
		return rotate(this)
	}
	this.right = insertKeyAndValue(this.right, key, value)
	this.height = maxHeight(this.left, this.right) + 1
	this.right.parent = this
	return rotate(this)

}
func (this *TreeMap) Insert(key string, value string) {
	this.root = insertKeyAndValue(this.root, key, value)
	return
}
func finMinNode(root *node) *node {

	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}
func finMaxNode(root *node) *node {

	if root == nil {
		return nil
	}
	for root.right != nil {
		root = root.right
	}
	return root
}
func findNode(r *node, key string) *node {
	//
	if r == nil {
		return nil
	}
	if strings.Compare(r.key, key) == 0 {
		return r
	}
	if strings.Compare(r.key, key) > 0 {
		return findNode(r.left, key)
	}
	return findNode(r.right, key)
}

func deleteNode(this *node, key string) *node {
	this = findNode(this, key)
	if this == nil {
		return nil
	}

	// strings.Compare(key, this.key) == 0
	if this.left == nil && this.right == nil {
		if this.parent == nil {
			return nil //根节点情况
		}
		if this.parent.left == this { //叶节点情况
			this.parent.left = nil
		} else {
			this.parent.right = nil
		}
		return this.parent //返回被删除节点的父节点
	}
	targetNode := finMinNode(this.right)
	if targetNode != nil {
		this.key = targetNode.key
		this.value = targetNode.value
		return deleteNode(targetNode, targetNode.key)
	}
	targetNode = finMaxNode(this.left)
	this.key = targetNode.key
	this.value = targetNode.value
	return deleteNode(targetNode, targetNode.key)
}
func adjust_tree(parent *node) *node { //删除节点的后检查操作，检查被节点父节点是否平衡，一直到根节点
	if parent == nil {
		return nil
	}
	for ; parent.parent != nil; parent = parent.parent {
		parent.height = maxHeight(parent.left, parent.right) + 1
		parent = rotate(parent)
	}
	return parent
}
func (this *TreeMap) GetValue(key string) string {
	node := findNode(this.root, key)
	if node == nil {
		return ""
	}
	return node.value
}
func (this *TreeMap) Delete(key string) {
	if findNode(this.root, key) == nil {
		return
	}
	parent := deleteNode(this.root, key)
	this.root = adjust_tree(parent)

}
func (this *TreeMap) ShowAll() {
	if this.root == nil {
		return
	}
	LeftTree := TreeMap{root: this.root.left}
	rightTree := TreeMap{root: this.root.right}
	LeftTree.ShowAll()
	fmt.Println(this.root.key + " : " + this.root.value)
	rightTree.ShowAll()
}

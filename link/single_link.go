package link

import "fmt"

type ElementInt int

//注意当前项目所有的链表都是不带表头结点的
type LinkNode struct {
	Node  ElementInt
	pNext *LinkNode
}

func (n *LinkNode) GetLen() int {
	if n == nil {
		return 0
	}
	curNode := n
	iLen := 0
	for curNode != nil {
		iLen++
		curNode = curNode.pNext
	}
	return iLen

}

func (n *LinkNode) PrintNode() {
	curNode := n
	for curNode != nil {
		fmt.Printf("%v ", curNode.Node)
		curNode = curNode.pNext
	}
	fmt.Println()
}

type LinkList struct {
	pHead *LinkNode
}

func CreateList(array []ElementInt) *LinkList {
	if len(array) == 0 {
		return nil
	}
	var retList LinkList
	var cur *LinkNode
	for _, v := range array {
		var node LinkNode
		node.Node = v
		if retList.pHead == nil {
			retList.pHead = &node
			cur = &node
		} else {
			cur.pNext = &node
			cur = cur.pNext
		}
	}

	return &retList
}

func (l *LinkList) GetHeadNode() *LinkNode {
	return l.pHead
}

//插入到第k个位置
func (l *LinkList) InsertNode(pos int, ele ElementInt) bool {
	iCount := 1 //注意这里应该定义为1
	curNode := l.pHead
	for curNode != nil && iCount != pos-1 {
		curNode = curNode.pNext
		iCount++
	}
	//元素不够pos-1个
	if iCount != pos-1 {
		return false
	}

	var node LinkNode
	node.Node = ele

	if curNode.pNext == nil {
		curNode.pNext = &node
		return true
	}

	node.pNext = curNode.pNext
	curNode.pNext = &node

	return true
}

//删除链表第K个节点
func (l *LinkList) DeleteNode(pos int) (ElementInt, bool) {
	iCount := 1
	cur := l.pHead
	for cur.pNext != nil && iCount != pos-1 {
		cur = cur.pNext
		iCount++
	}
	if iCount != pos-1 {
		fmt.Printf("link node not exist \n")
		return -1, false
	}

	//没有第pos个节点
	if cur.pNext == nil {
		fmt.Printf("link node not exist \n")
		return -1, false
	}
	retNode := cur.pNext
	cur.pNext = cur.pNext.pNext
	return retNode.Node, true
}

//查找链表中是否存在某个节点x,如果有则删除该节点并返回false，如果没有则返回true
func (l *LinkList) FindAndDelete(x ElementInt) bool {
	if l.pHead == nil {
		return false
	}

	curNode := l.pHead
	for curNode != nil && curNode.Node != x {
		curNode = curNode.pNext
	}
	if curNode == nil {
		return false
	}
	return true
}
func (l *LinkList) Print() {
	curNode := l.pHead
	for curNode != nil {
		fmt.Printf("%v ", curNode.Node)
		curNode = curNode.pNext
	}
	fmt.Println()

}

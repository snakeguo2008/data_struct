package link

import "fmt"

type DoubleLink struct {
	Ele   ElementInt
	prior *DoubleLink
	next  *DoubleLink
}

func (d *DoubleLink) PrintDB() {
	if d == nil {
		return
	}
	for d != nil {
		fmt.Printf("%v ", d.Ele)
	}
	fmt.Println()
}

//创建一个单链表
func CreateDBLink(array []ElementInt) *DoubleLink {
	if len(array) == 0 {
		return nil
	}

	var pHead, curNode *DoubleLink
	for _, v := range array {
		var Node DoubleLink
		Node.Ele = v
		if pHead == nil {
			pHead = &Node
			curNode = &Node
		} else {
			curNode.next = &Node
			curNode = curNode.next
		}
	}
	return pHead

}

//查找指定元素的节点
func (d *DoubleLink) FindNode(x ElementInt) *DoubleLink {
	if d == nil {
		return nil
	}
	curNode := d
	for curNode != nil && curNode.Ele != x {
		curNode = curNode.next
	}
	return curNode
}

//

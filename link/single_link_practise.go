package link

import "fmt"

/***************
****合并两个非递减有序的链表A和B
****模式一:循环模式
****特点:代码比较易懂，但是比较冗余
****当前的链表不带表头结点
****************/

func MergeLink(LinkA, LinkB *LinkNode) *LinkNode {
	if LinkA == nil {
		return LinkB
	} else if LinkB == nil {
		return LinkA
	}

	var pHead, pLinkNode *LinkNode

	for LinkA != nil && LinkB != nil {
		var curNode LinkNode
		if LinkA.Node > LinkB.Node {

			curNode.Node = LinkB.Node
			LinkB = LinkB.pNext

		} else {
			curNode.Node = LinkA.Node
			LinkA = LinkA.pNext

		}
		//链表从0~1的时候
		if pHead == nil {
			pHead = &curNode
			pLinkNode = &curNode
		} else {
			pLinkNode.pNext = &curNode
			pLinkNode = pLinkNode.pNext
		}

	}

	if LinkA != nil {
		pLinkNode.pNext = LinkA
	} else if LinkB != nil {
		pLinkNode.pNext = LinkB

	}

	return pHead
}

/***************
****合并两个非递减有序的链表A和B
****模式而: 递归模式
****特点: 代码简洁，但是不易理解
****当前的链表不带表头结点
****************/
//做递归有返回的题材需要注意的是每个递归的return都会被上层递归捕获
func MergeLinkV2(LinkA, LinkB *LinkNode) *LinkNode {
	if LinkA == nil {
		return LinkB
	} else if LinkB == nil {
		return LinkA
	}
	//将它视为递归方法中返回的节点
	var retNode *LinkNode

	if LinkA.Node > LinkB.Node {
		//那么该节点应该返回更小的那个
		retNode = LinkB
		retNode.pNext = MergeLinkV2(LinkA, LinkB.pNext)
	} else {
		retNode = LinkA
		retNode.pNext = MergeLinkV2(LinkA.pNext, LinkB)
	}
	return retNode
}

/****##############单链表逆置#################**/

/***********************
** 1)定义返回头指针
   2)拆除当前节点(注意保留下一个节点的位置)
** 3)当前节点指向返回头指针指向的节点
** 4)移动返回头指针
*/
func ReverseList(pHead *LinkNode) *LinkNode {
	if pHead == nil {
		return nil
	}
	//定义头指针
	var retHeadNode *LinkNode = nil
	for pHead != nil {
		cur := pHead
		//保存下个节点
		pHead = cur.pNext
		//倒插法
		cur.pNext = retHeadNode
		retHeadNode = cur
	}
	return retHeadNode
}

/***********************
***	找链表倒数第K个节点
*** 思路：先定义一个节点，走出K步，然后再定义一个节点同步走
*** 特别注意：如果是应试,需要特别注意代码的健壮性！！！！
*****/
func FindPosReverse(pNode *LinkNode, k int) *LinkNode {
	if k <= 0 || pNode == nil {
		return nil
	}
	var stepFront, stepBack = pNode, pNode
	step := 1

	for stepFront != nil && step != k {
		stepFront = stepFront.pNext
		step++
	}
	if step != k {
		return nil
	}

	//注意这里的处理细节
	for stepFront.pNext != nil {
		stepFront = stepFront.pNext
		stepBack = stepBack.pNext

	}
	return stepBack
}

/************找出两个链表的交汇点**************/
/***********************
***思路1: 如果两个链表有交汇，根据链表的特征，那么交汇节点之后所有的节点都相同。可以将两个
***      链表都都入栈，从后往前开始比较，找到第一个不相同的节点，那么上个出栈节点就是公共节点
***思路2: 如果两个链表交汇，那么交汇之后的长度相同，但是交汇之前的长度不同。那么先获取两个链表的长度
***		 ，长(A)的先移动直至剩下链表长度和B相同为止。然后在同时移动。那么第一个公共节点即为所找
***/
func FindCommNode(pNodeA *LinkNode, pNodeB *LinkNode) *LinkNode {
	if pNodeA == nil || pNodeB == nil {
		return nil
	}

	iLenA := pNodeA.GetLen()
	iLenB := pNodeB.GetLen()
	step := iLenA - iLenB

	if step > 0 {
		for i := 0; i < step; i++ {
			pNodeA = pNodeA.pNext
		}
	} else {
		for i := 0; i < step; i++ {
			pNodeB = pNodeB.pNext
		}
	}

	//移动step步长之后，剩下的链长度相同.可以用其中一链作为循坏条件
	for pNodeA != nil {
		if pNodeA == pNodeB {
			break
		}
		pNodeA = pNodeA.pNext
		pNodeB = pNodeB.pNext

	}
	return pNodeA
}

/******逆向打印链表元素************
** 1) 逆向首先想到栈，然后想到递归
**
**
**/
func ReversePrint(pNodeA *LinkNode) {
	if pNodeA != nil {

		ReversePrint(pNodeA.pNext)
		fmt.Printf("%v ", pNodeA.Node)
	}
}

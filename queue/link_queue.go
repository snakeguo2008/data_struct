package queue

//链队列

type QNode struct {
	front *QNode
	rear  *QNode
}

func InitLinkQueue() *QNode {
	return &QNode{
		front: nil,
		rear:  nil,
	}
}

//判断队空
func (s *QNode) IsEmpty() bool {
	if s.front == nil || s.rear == nil {
		return true
	}
	return false
}

package queue

//循环队列
type StQueue struct {
	data    []int
	front   int //循环队列中指向的是队首前一个空的元素格
	rear    int //指向的是队列尾部的元素个
	maxSize int
}

func NewQueue(size_t int) *StQueue {
	//队首和队尾重合表示队列为空，然后初始化为0
	return &StQueue{
		data:    make([]int, size_t),
		front:   0,
		rear:    0,
		maxSize: size_t,
	}
}

func (s *StQueue) IsEmpty() bool {
	return s.front == s.rear
}

func (s *StQueue) PushQueue(x int) bool {
	//判断队列是否为满,循环队列留了一个位置表示位置是否满了。
	if (s.rear+1)%s.maxSize == s.front {
		return false
	}
	//若未满则需要移动指针，这里需要注意循环对了的移动方式
	s.rear = (s.rear + 1) % s.maxSize
	s.data[s.rear] = x
	return true
}

func (s *StQueue) PopQueue() (int, bool) {
	//判断是否为空
	if s.front == s.rear {
		return -1, false
	}
	s.front = (s.front + 1) % s.maxSize
	return s.data[s.front], true
}

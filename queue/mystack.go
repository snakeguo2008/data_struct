package queue

type SqStack struct {
	data    []int
	top     int //当栈为空时,top = -1
	maxSize int
}

func NewSqStack(MaxSize int) *SqStack {
	return &SqStack{
		data:    make([]int, MaxSize),
		top:     -1,
		maxSize: MaxSize,
	}
}

func (s *SqStack) IsEmpty() bool {
	return s.top == -1
}

func (s *SqStack) Push(x int) bool {
	//每次进栈判断是否栈满
	if s.top == s.maxSize {
		return false
	}

	s.data[s.top] = x
	s.top++
	return true
}

func (s *SqStack) Pop() (int, bool) {
	//判断栈空
	if s.top == -1 {
		return -1, false
	}
	x := s.data[s.top]
	s.top--
	return x, false
}

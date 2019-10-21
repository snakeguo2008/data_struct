package queue

/*
利用两个栈模拟一个队列
解题: 队列的特征是先进先出， 栈的特性是先进后出。
*/

type TStack struct {
	element []int
	top     [2]int
}

func InitTStack(maxSize int) *TStack {
	return &TStack{
		element: make([]int, maxSize),
		top:     [2]int{-1, -1},
	}
}

//用空间共享栈实现入栈出栈操作
//stNO是栈的编号，表示进入哪个栈
func (s *TStack) SharePush(x, stNo int) bool {
	//判断栈满
	if s.top[0]+1 >= s.top[1] {
		return false
	}

	if stNo == 0 {
		s.top[0]++
		s.element[s.top[0]] = x
		return true
	} else {
		s.top[1]++
		s.top[1] = x
		return true
	}
	return false
}

func (s *TStack) SharePop(stNo int) int {
	if stNo == 0 {
		//判断栈0是否为空
		if s.top[0] == -1 {
			return -1
		}
		x := s.element[s.top[0]]
		s.top[0]--
		return x
	} else if stNo == 1 {
		if s.top[1] == -1 {
			return -1
		}
		x := s.element[s.top[1]]
		s.top[1]--
		return x
	}
	return -1
}

/*****
用栈来模拟对
************/

package link

import "fmt"

type Element interface{}

const (
	DefaultSize = 128
)

/***
顺序表
1、当前容量
2、当前校验
3、查找
4、插入
5、删除
****/
type SqList struct {
	sqList  []Element
	iLength int //当前顺序表的长度
}

func NewSqList(length int) *SqList {
	if length <= 0 {
		length = DefaultSize
	}
	return &SqList{make([]Element, length, length), 0}
}

//容量
func (s *SqList) CapSqList() int {
	return cap(s.sqList)
}

/*插入到位置pos处*/
/*[0 ~ cap(s.sqList))]*/
func (s *SqList) InsertElem(element Element, pos int) error {
	if err := s.checkInsertPos(pos, true); err != nil {
		return err
	}

	for i := s.iLength; i > pos; i-- {
		s.sqList[i+1] = s.sqList[i]
	}
	s.sqList[pos] = element
	s.iLength++

	return nil
}

//删除位置为pos的元素并返回其值
func (s *SqList) DeleteElem(pos int) (Element, error) {
	if err := s.checkInsertPos(pos, false); err != nil {
		return nil, err
	}
	ret := s.sqList[pos-1]
	for i := pos; i < s.iLength; i++ {
		s.sqList[i-1] = s.sqList[i+1]
	}
	return ret, nil
}

/*查找元素位置*/
func (s *SqList) FindElem(element Element) bool {
	for _, x := range s.sqList {
		if x == element {
			return true
		}
	}
	return false
}

//添加代码
func (s *SqList) AppendElem(element Element) bool {
	if s.iLength >= s.CapSqList() {
		return false
	}
	fmt.Println(s.sqList)

	s.sqList[s.iLength] = element
	s.iLength++
	return true
}

/**/
func (s *SqList) PrintSqList() {
	for ele := range s.sqList {
		fmt.Println(ele)
	}
}

func (s *SqList) PrintCurLen() {
	fmt.Println(s.iLength)
}

//那么position的范围是0~n-1
func (s *SqList) checkInsertPos(pos int, bInsert bool) error {
	//1、判断是否插入的边界之外;
	if pos < 0 || pos > s.iLength {
		return fmt.Errorf("out of range")
	}

	//2、判断插入的数组是否已经满了
	if bInsert && s.iLength == s.CapSqList() {
		return fmt.Errorf("the sqList is full!")
	}

	return nil
}

package data

type Node struct {
	value float64
	next  *Node
}

//栈，先进后出，后进先出
type MyStack struct {
	head *Node
	size int
}

//判断栈是否为空
func (st *MyStack) Empty() bool {
	if st.size > 0 {
		return false
	}
	return true
}

//获取栈的元素个数
func (st *MyStack) Size() int {
	return st.size
}

//入栈
func (st *MyStack) Push(v float64) {
	//使用头插法
	newHead := new(Node)
	newHead.value = v
	newHead.next = st.head
	st.head = newHead
	st.size += 1
}

//出栈
func (st *MyStack) Pop() {
	if st.size > 0 {
		next := st.head.next
		st.head.next = nil
		st.head = next
		st.size -= 1
	}
}

//获取栈顶元素
func (st *MyStack) Top() float64 {
	if st.size > 0 {
		return st.head.value
	}
	return 0
}

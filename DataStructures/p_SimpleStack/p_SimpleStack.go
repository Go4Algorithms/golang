package p_SimpleStack

type SimpleStack struct {
  arr []int
  top int
}

func New() SimpleStack {
  ss := SimpleStack{}
  ss.arr = make([]int, 5)
  ss.top = 0
  return ss
}

func (sq *SimpleStack) IsEmpty() bool {
  return top == 0
}

func (ss *SimpleStack) Push(num int) {
  if (ss.top >= len(ss.arr) - 1) {
    ss.arr = append(ss.arr, make([]int, 5)...)
  }
  ss.top++
  ss.arr[ss.top] = num
}

func (ss *SimpleStack) Pop() int {
  if (ss.top > 0) {
    ss.top--
    return ss.arr[ss.top + 1]
  }
  return 0
}

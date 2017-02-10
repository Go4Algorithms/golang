package p_SimpleQueue

type SimpleQueue struct {
  arr []int
  top int
}

func New() SimpleQueue {
	sq := SimpleQueue{}
  sq.arr = make([]int, 5)
  sq.top = 0
  return sq
}

func (sq *SimpleQueue) Push(num int) {
  if (sq.top < len(sq.arr)) {
    sq.arr = append(sq.arr, make([]int, 5)...)
  }
  sq.top++
  sq.arr[sq.top] = num
}

func (sq *SimpleQueue) Pop() int {
  if (sq.top > 0) {
		for i := 0; i < len(sq.arr) - 1; i++ {
			sq.arr[i] = sq.arr[i + 1]
		}
    sq.top--
    return sq.arr[0]
  }
  return 0
}


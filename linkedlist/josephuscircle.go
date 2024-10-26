package linkedlist

type JosephusCircle struct {
	N      int
	M      int
	Result int
}

func NewJosephusCircle(N, M int) *JosephusCircle {
	return &JosephusCircle{
		N: N,
		M: M,
	}
}

func (jc *JosephusCircle) GetJosephusPosition() int {
	val := []int{}
	for i := 1; i <= jc.N; i++ {
		val = append(val, i)
	}
	head := ConvertArrayToCircularLinkedList(val)
	p := head
	for count := jc.N; count > 1; count-- {
		for i := 0; i < jc.M-1; i++ {
			p = p.GetNext()
		}
		p.SetNext(p.GetNext().GetNext())
	}
	return p.GetData()
}

func (jc *JosephusCircle) Execute() {
	jc.Result = jc.GetJosephusPosition()
}

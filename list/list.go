package list

type ListNode struct {
	str string
	next *ListNode
	mutex *MyMutex
}

type List struct {
	first *ListNode
	size int
	mutex *MyMutex
}

func NewList() *List {
	list := List{nil, 0, &MyMutex{}}

	first := createNode("")
	list.first = first

	return &list
}

func createNode(str string) *ListNode {
	return &ListNode{str, nil, &MyMutex{}}
}

func (l *List) AddToBegin (str string) {
	node := createNode(str)

	l.first.mutex.Lock()
	node.next = l.first.next
	l.first.next = node
	l.first.mutex.Unlock()

	l.mutex.Lock()
	l.size++
	l.mutex.Unlock()
}

func (l *List) isEmpty() bool {
	return l.size == 0
}

func (l *List) String() string {
	if (l.isEmpty()) {
		return "Empty\n"
	}

	str := ""

	it := l.first
	var prev *ListNode
	it.mutex.Lock()

	for it != nil {
		if (it.next != nil) {
			if (len(it.str) != 0) {
				str += it.str
				str += " -> "
			}
		} else {
			if (len(it.str) != 0) {
				str += it.str 
			}
		}

		prev = it
		it = it.next

		if (it != nil) {
			it.mutex.Lock()
		}
		
		if prev.mutex.IsLocked() {
			prev.mutex.Unlock()
		}
	}

	return str
}

func swap(first *ListNode, second *ListNode) {
	var temp *ListNode
	elem1 := first.next
	elem2 := second.next
	first.next = elem2
	second.next = elem1
	temp = elem1.next
	elem1.next = elem2.next
	elem2.next = temp 
}

func (l *List) Sort() {
	var size int
	
	l.mutex.Lock()
	size = l.size
	l.mutex.Unlock()

	var iEl, jEl, iPrev, jPrev *ListNode
	var isNear bool

	if size <= 1 {
		return
	}

	iEl = l.first
	iEl.mutex.Lock()

	for i := 0; i < size; i++ {
		jEl = iEl.next

		jEl.mutex.Lock()

		for j := i + 1; j < size; j++ {
			jEl.next.mutex.Lock()

			if iEl.next != jEl {
				iEl.next.mutex.Lock()
			}

			if iEl.next.str > jEl.next.str {
				isNear = iEl.next == jEl
				
				swap(iEl, jEl)
			
				if (isNear) {
					jEl = iEl.next
				}
			}
			
			jPrev = jEl
			jEl = jEl.next

			if iEl.next != jEl {
				iEl.next.mutex.Unlock()
			}

			if jPrev.mutex.IsLocked() {
				jPrev.mutex.Unlock()
			}
		}

		jEl.mutex.Unlock()
		iEl.next.mutex.Lock()

		iPrev = iEl
		iEl = iEl.next

		iPrev.mutex.Unlock()
	}

	iEl.mutex.Unlock()
}




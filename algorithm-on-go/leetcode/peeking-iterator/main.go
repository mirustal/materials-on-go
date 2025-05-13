package peeking_iterator

type PeekingIterator struct {
	iterator    *Iterator
	peekElement int
}

func Constructor(iter *Iterator) *PeekingIterator {
	pi := &PeekingIterator{iter, 0}
	if iter.hasNext() {
		pi.peekElement = iter.next()
	}

	return pi
}

func (this *PeekingIterator) next() int {
	res := this.peekElement
	if this.iterator.hasNext() {
		this.peekElement = this.iterator.next()
	} else {
		this.peekElement = 0
	}
	return res
}

func (this *PeekingIterator) hasNext() bool {
	return this.peekElement != 0
}

func (this *PeekingIterator) peek() int {
	return this.peekElement
}

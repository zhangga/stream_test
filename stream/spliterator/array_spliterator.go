package spliterator

import "github.com/zhangga/stream_test/stream/types"

type ArraySpliterator struct {

	// array
	Array []types.T

	// current index, modified on advance/split
	Index int
	// one past last index
	Fence int
	// characteristics
	Characteristics int
}

func (spliterator *ArraySpliterator) tryAdvance(action types.Consumer) bool {
	if action == nil {
		return false
	}
	if spliterator.Index >= 0 && spliterator.Index < spliterator.Fence {
		e := spliterator.Array[spliterator.Index]
		spliterator.Index++
		action(e)
		return true
	}
	return false
}

func (spliterator *ArraySpliterator) forEachRemaining(action types.Consumer) {
	if action == nil {
		return
	}
}

func (spliterator *ArraySpliterator) getCharacteristics() int {
	return spliterator.Characteristics
}

func (spliterator *ArraySpliterator) getComparator() {

}

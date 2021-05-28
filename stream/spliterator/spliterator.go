package spliterator

import "github.com/zhangga/stream_test/stream/types"

const (
	ORDERED		= 0x00000010
	DISTINCT	= 0x00000001
	SORTED		= 0x00000004
	SIZED		= 0x00000040
	NONNULL		= 0x00000100
	IMMUTABLE	= 0x00000400
	CONCURRENT	= 0x00001000
	SUBSIZED	= 0x00004000
)

type Spliterator interface {

	// tryAdvance If a remaining element exists, performs the given action on it,
	tryAdvance(action types.Consumer) bool

	forEachRemaining(action types.Consumer)

	getCharacteristics() int

	getComparator()

}

func OpFlagFromCharacteristics(spliterator Spliterator) {
	if (spliterator.getCharacteristics() & SORTED) != 0 && spliterator.getComparator() != nil) {
		// Do not propagate the SORTED characteristic if it does not correspond
		// to a natural sort order
		return spliterator.getCharacteristics() & SPLITERATOR_CHARACTERISTICS_MASK & ~SORTED
	} else {
		return spliterator.getCharacteristics() & SPLITERATOR_CHARACTERISTICS_MASK
	}

}

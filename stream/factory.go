package stream

import (
	"github.com/zhangga/stream_test/stream/spliterator"
	"github.com/zhangga/stream_test/stream/types"
)

func Of(elements ...types.T) Stream {
	spliterator := &spliterator.ArraySpliterator{
		Array:           elements,
		Index:           0,
		Fence:           len(elements),
		Characteristics: spliterator.ORDERED | spliterator.IMMUTABLE | spliterator.SIZED | spliterator.SUBSIZED,
	}
	return stream(spliterator, false)
}

func stream(spliterator spliterator.Spliterator, parallel bool) Stream {

}

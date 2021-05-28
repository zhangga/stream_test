package stream

import "github.com/zhangga/stream_test/stream/types"

type Stream interface {

	// Filter Returns a stream consisting of the elements of this stream that match the given predicate.
	Filter(types.Predicate) Stream
}

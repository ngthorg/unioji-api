package connections

import (
	"math"

	"github.com/unioji/unioji-api/graph/relay"
)

// HPageInfo struct
type HPageInfo struct {
	hasNext     bool
	hasPrev     bool
	startOffset int
}

// HandlinPageInfo func
func HandlinPageInfo(args relay.ConnectionArgs, meta relay.SliceMetaInfo) HPageInfo {
	sliceStart := 0
	sliceEnd := sliceStart + meta.Length
	beforeOffset := meta.Length
	afterOffset := -1
	lowerBound := 0
	upperBound := meta.Length
	hasNextPage := false
	hasPreviousPage := false
	if args.Before != nil {
		beforeOffset, _ = relay.CursorToOffset(*args.Before)
		upperBound = beforeOffset
	}
	if args.After != nil {
		afterOffset, _ = relay.CursorToOffset(*args.After)
		lowerBound = afterOffset + 1
	}
	startOffset := int(math.Max(float64(sliceStart-1), math.Max(float64(afterOffset), -1))) + 1
	endOffset := int(math.Min(float64(sliceEnd), math.Min(float64(beforeOffset), float64(meta.Length))))
	if args.First != nil {
		endOffset = int(math.Min(float64(endOffset), float64(startOffset+*args.First)))
		hasNextPage = endOffset < upperBound
	}
	if args.Last != nil {
		startOffset = int(math.Max(float64(startOffset), float64(endOffset-*args.Last)))
		hasPreviousPage = startOffset > lowerBound
	}
	return HPageInfo{
		hasNext:     hasNextPage,
		hasPrev:     hasPreviousPage,
		startOffset: startOffset,
	}
}

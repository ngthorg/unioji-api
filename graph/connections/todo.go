package connections

import (
	"errors"
	"math"

	"github.com/unioji/unioji-api/graph/model"
	"github.com/unioji/unioji-api/graph/relay"
)

// ConnectionFromSliceTodos return TodoConnection
func ConnectionFromSliceTodos(todos []*model.Todo, args relay.ConnectionArgs,
	meta relay.SliceMetaInfo,
) (*model.TodoConnection, error) {
	if args.First != nil && args.Last != nil {
		return nil, errors.New(`
			Passing both first and last to paginatethe connection is not supported.
		`)

	}

	if args.Before != nil && args.After != nil {
		return nil, errors.New(`
			Passing both before and after to paginate the connection is not supported.
		`)
	}

	if args.First != nil {
		if *args.First == 0 {
			return &model.TodoConnection{
				Edges: []*model.TodoEdge{},
				Nodes: []*model.Todo{},
			}, nil
		} else if *args.First < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

	if args.Last != nil {
		if *args.Last == 0 {
			return &model.TodoConnection{
				Edges: []*model.TodoEdge{},
				Nodes: []*model.Todo{},
			}, nil
		} else if *args.Last < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

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

	edges := []*model.TodoEdge{}
	for i, todo := range todos {
		edges = append(edges, &model.TodoEdge{
			Node:   todo,
			Cursor: relay.OffsetToCursor(startOffset + i),
		})
	}

	var startCursor, endCursor string
	if len(edges) > 0 {
		startCursor = edges[0].Cursor
		endCursor = edges[len(edges)-1].Cursor
	}

	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     hasNextPage,
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
		},
		// Edges:      edges,
		Nodes:      todos,
		TotalCount: meta.Length,
	}, nil
}

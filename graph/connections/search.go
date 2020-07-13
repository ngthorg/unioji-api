package connections

import (
	"errors"

	"github.com/unioji/unioji-api/graph/model"
	"github.com/unioji/unioji-api/graph/relay"
)

// ConnectionFromSliceSearches return NodeConnection
func ConnectionFromSliceSearches(searches []model.SearchResult, args relay.ConnectionArgs,
	meta relay.SliceMetaInfo,
) (*model.SearchResultConnection, error) {
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
			return &model.SearchResultConnection{
				Edges: []*model.SearchResultEdge{},
				Nodes: []model.SearchResult{},
			}, nil
		} else if *args.First < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

	if args.Last != nil {
		if *args.Last == 0 {
			return &model.SearchResultConnection{
				Edges: []*model.SearchResultEdge{},
				Nodes: []model.SearchResult{},
			}, nil
		} else if *args.Last < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

	pageInfo := HandlinPageInfo(args, meta)

	edges := []*model.SearchResultEdge{}
	for i, search := range searches {
		edges = append(edges, &model.SearchResultEdge{
			Node:   search,
			Cursor: relay.OffsetToCursor(pageInfo.startOffset + i),
		})
	}

	var startCursor, endCursor string
	if len(edges) > 0 {
		startCursor = edges[0].Cursor
		endCursor = edges[len(edges)-1].Cursor
	}

	return &model.SearchResultConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: pageInfo.hasPrev,
			HasNextPage:     pageInfo.hasNext,
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
		},
		Edges:      edges,
		Nodes:      searches,
		TotalCount: meta.Length,
	}, nil
}

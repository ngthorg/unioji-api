package connections

import (
	"errors"

	"github.com/unioji/unioji-api/graph/model"
	"github.com/unioji/unioji-api/graph/relay"
)

// ConnectionFromSliceNodes return NodeConnection
func ConnectionFromSliceNodes(nodes []model.Node, args relay.ConnectionArgs,
	meta relay.SliceMetaInfo,
) (*model.NodeConnection, error) {
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
			return &model.NodeConnection{
				Edges: []*model.NodeEdge{},
				Nodes: []model.Node{},
			}, nil
		} else if *args.First < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

	if args.Last != nil {
		if *args.Last == 0 {
			return &model.NodeConnection{
				Edges: []*model.NodeEdge{},
				Nodes: []model.Node{},
			}, nil
		} else if *args.Last < 0 {
			return nil, relay.ErrInvalidPagination
		}
	}

	pageInfo := HandlinPageInfo(args, meta)

	edges := []*model.NodeEdge{}
	for i, node := range nodes {
		edges = append(edges, &model.NodeEdge{
			Node:   node,
			Cursor: relay.OffsetToCursor(pageInfo.startOffset + i),
		})
	}

	var startCursor, endCursor string
	if len(edges) > 0 {
		startCursor = edges[0].Cursor
		endCursor = edges[len(edges)-1].Cursor
	}

	return &model.NodeConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: pageInfo.hasPrev,
			HasNextPage:     pageInfo.hasNext,
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
		},
		Edges:      edges,
		Nodes:      nodes,
		TotalCount: meta.Length,
	}, nil
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type SearchResult interface {
	IsSearchResult()
}

type NodeConnection struct {
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*NodeEdge `json:"edges"`
	Nodes      []Node      `json:"nodes"`
	TotalCount int         `json:"totalCount"`
}

type NodeEdge struct {
	Node   Node   `json:"node"`
	Cursor string `json:"cursor"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type TodoConnection struct {
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*TodoEdge `json:"edges"`
	Nodes      []*Todo     `json:"nodes"`
	TotalCount int         `json:"totalCount"`
}

type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

type TodoOrderBy string

const (
	TodoOrderByCreatedAtAsc  TodoOrderBy = "createdAt_ASC"
	TodoOrderByCreatedAtDesc TodoOrderBy = "createdAt_DESC"
	TodoOrderByUpdatedAtAsc  TodoOrderBy = "updatedAt_ASC"
	TodoOrderByUpdatedAtDesc TodoOrderBy = "updatedAt_DESC"
)

var AllTodoOrderBy = []TodoOrderBy{
	TodoOrderByCreatedAtAsc,
	TodoOrderByCreatedAtDesc,
	TodoOrderByUpdatedAtAsc,
	TodoOrderByUpdatedAtDesc,
}

func (e TodoOrderBy) IsValid() bool {
	switch e {
	case TodoOrderByCreatedAtAsc, TodoOrderByCreatedAtDesc, TodoOrderByUpdatedAtAsc, TodoOrderByUpdatedAtDesc:
		return true
	}
	return false
}

func (e TodoOrderBy) String() string {
	return string(e)
}

func (e *TodoOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoOrderBy", str)
	}
	return nil
}

func (e TodoOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

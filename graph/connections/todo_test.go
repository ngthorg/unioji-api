package connections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unioji/unioji-api/graph/model"
	"github.com/unioji/unioji-api/graph/relay"
)

type Input struct {
	before *string
	after  *string
	first  *int
	last   *int
	lenght int
}

type Expected struct {
	hasNext     bool
	hasPrev     bool
	startCursor string
	endCursor   string
	nodes       []*model.Todo
	total       int
}

func TestConnectionFromSliceTodos(t *testing.T) {
	assert := assert.New(t)
	todos := []*model.Todo{
		{
			ID:        "1",
			Text:      "Test 1",
			Completed: true,
		},
		{
			ID:        "2",
			Text:      "Demo",
			Completed: false,
		},
		{
			ID:        "3",
			Text:      "Test 3",
			Completed: false,
		},
		{
			ID:        "4",
			Text:      "Test 4",
			Completed: false,
		},
		{
			ID:        "5",
			Text:      "Test 5",
			Completed: false,
		},
		{
			ID:        "6",
			Text:      "Test 6",
			Completed: false,
		},
		{
			ID:        "7",
			Text:      "Test 7",
			Completed: false,
		},
		{
			ID:        "8",
			Text:      "Test 8",
			Completed: false,
		},
		{
			ID:        "9",
			Text:      "Test 9",
			Completed: false,
		},
		{
			ID:        "10",
			Text:      "Test 10",
			Completed: false,
		},
	}

	f := 10
	l := 10
	var tests = []struct {
		input    Input
		expected Expected
	}{
		{
			input: Input{
				first:  &f,
				last:   nil,
				before: nil,
				after:  nil,
				lenght: 20,
			},
			expected: Expected{
				hasNext:     true,
				hasPrev:     false,
				startCursor: relay.OffsetToCursor(0),
				endCursor:   relay.OffsetToCursor(9),
				nodes:       todos,
				total:       20,
			},
		},
		{
			input: Input{
				first:  nil,
				last:   &l,
				before: nil,
				after:  nil,
				lenght: 20,
			},
			expected: Expected{
				hasNext:     false,
				hasPrev:     true,
				startCursor: relay.OffsetToCursor(10),
				endCursor:   relay.OffsetToCursor(19),
				nodes:       todos,
				total:       20,
			},
		},
	}

	for _, test := range tests {
		result, _ := ConnectionFromSliceTodos(todos, relay.ConnectionArgs{
			Before: test.input.before,
			After:  test.input.after,
			First:  test.input.first,
			Last:   test.input.last,
		}, relay.SliceMetaInfo{
			Length: test.input.lenght,
		})
		assert.Equal(&model.TodoConnection{
			PageInfo: &model.PageInfo{
				HasPreviousPage: test.expected.hasPrev,
				HasNextPage:     test.expected.hasNext,
				StartCursor:     &test.expected.startCursor,
				EndCursor:       &test.expected.endCursor,
			},
			// Edges:      edges,
			Nodes:      todos,
			TotalCount: test.expected.total,
		}, result, fmt.Sprintf("they should equal %v", test.input))
	}
}

package relay

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrInvalidPagination error
var ErrInvalidPagination = errors.New("ent: invalid pagination parameters")

// ConnectionArgs params default in a relay connection
type ConnectionArgs struct {
	Before *string
	After  *string
	First  *int
	Last   *int
}

// SliceMetaInfo total of all
type SliceMetaInfo struct {
	Length int
}

const cursorPrefix = "cursor:"

// Edge define the contract for an edge in a relay connection
type Edge interface {
	GetCursor() string
}

// OffsetToCursor create the cursor string from an offset
func OffsetToCursor(offset int) string {
	str := fmt.Sprintf("%v%v", cursorPrefix, offset)
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// CursorToOffset re-derives the offset from the cursor string.
func CursorToOffset(cursor string) (int, error) {
	str := ""
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err == nil {
		str = string(b)
	}
	str = strings.Replace(str, cursorPrefix, "", -1)
	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("Invalid cursor")
	}
	return offset, nil
}

package reverse

import (
	"strings"

	"github.com/clipperhouse/uax29/graphemes"
)

func String(s string) string {
	segments := graphemes.SegmentAll([]byte(s))

	for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
		segments[i], segments[j] = segments[j], segments[i]
	}

	var slice []string

	for i := 0; i < len(segments); i++ {
		slice = append(slice, string(segments[i]))
	}

	return strings.Join(slice, "")
}

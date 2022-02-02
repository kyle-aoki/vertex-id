package vertexId

import (
	"fmt"
	"testing"
)

func TestVertexId(t *testing.T) {
	vid := VertexId()
	fmt.Println(vid)
	if len(vid) != 64 {
		t.Fail()
	}
}

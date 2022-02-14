package json_spanner

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	b := NewBuilder()
	b.Set("a", 6)
	b.Set("a.b", 3)
	b.Set("a.d", 3)

	b.Set("c", "xxx")
	fmt.Println(b.ToJson())

}

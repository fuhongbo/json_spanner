package wildcard

import (
	"fmt"
	"testing"
)

func Test_wildcard(t *testing.T) {

	fmt.Printf("%v\n", Match("test/foo/bar", "test/foo/bar"))
	fmt.Printf("%v\n", Match("test/foo/bar", "test/+/bar"))
	fmt.Printf("%v\n", Match("test/foo/bar", "test/#"))
	fmt.Printf("%v\n", Match("test/foo/bar/baz", "test/+/#"))
	fmt.Printf("%v\n", Match("test/foo/bar/baz", "test/+/+/baz"))
	fmt.Printf("%v\n", Match("test", "test/#"))
	fmt.Printf("%v\n", Match("test/", "test/#"))
	fmt.Printf("%v\n", Match("test/foo/bar", "test/+"))
	fmt.Printf("%v\n", Match("test/foo/bar", "test/nope/bar"))

}

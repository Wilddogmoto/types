package maps_test

import (
	"fmt"
	"github.com/Wilddogmoto/types/maps"
)

func ExampleMakeMap() {

	mp := maps.MakeMap[string, int](3)

	mp.Add("1", 1)
	mp.Add("2", 2)
	mp.Add("3", 3)

	fmt.Print(mp)
	// Output: map[1:1 2:2 3:3]
}

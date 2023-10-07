package main

import (
	"github.com/Wilddogmoto/types/maps"
	"github.com/Wilddogmoto/types/slices"
	"log"
)

func main() {

	sl := slices.MakeSlice[string](0, 0)

	sl.Append("Bob", "Greg", "Alice")

	sl.Sort(nil)

	log.Printf("slice: %+v", sl)

	mp := sl.Map()

	newmp := maps.Map[int, string].Copy(mp)

	log.Printf("map: %+v", newmp)

}

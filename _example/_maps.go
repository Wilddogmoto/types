package main

import (
	"github.com/Wilddogmoto/types/maps"
	"github.com/Wilddogmoto/types/slices"
	"log"
)

func main() {

	mp := maps.MakeMap[string, int](0)

	mp.Merge(map[string]int{"test": 100, "some": 11})

	sl := mp.Keys()

	//некоторые функции кастомного слайса Slice[Type] не умеет принимать параметры типа Slice[Type], но при создании пустого слайса, они способна принять слайс того же типа и выполнить свой функционал
	// example:
	//len := slices.Slice[string].Len(sl)
	//cap := slices.Slice[string].Cap(sl)
	//slice := slices.Slice[string].Async(sl)
	//max := slices.Slice[string].Max(sl)
	//min := slices.Slice[string].Min(sl)

	newslice := slices.Slice[string].Copy(sl)

	log.Printf("%+v", newslice)
}

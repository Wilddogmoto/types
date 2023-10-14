package channels_test

import (
	"fmt"
	"github.com/Wilddogmoto/types/channels"
	"sync"
)

func ExampleMakeChannel() {

	c := channels.MakeChannel[int](0)

	wg := sync.WaitGroup{}

	// writer
	wg.Add(1)
	go func() {
		defer func() {
			c.Close()
			wg.Done()
		}()
		c.Write(1)
		c.Write(2)
		c.Write(3)
	}()

	values := make([]int, 0, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Range(func(value int) {
			values = append(values, value)
		})
	}()

	wg.Wait()

	fmt.Print(values)
	// Output: [1 2 3]
}

func ExampleChannel_Merge() {

	chan1 := makeChannel[int]([]int{1, 2, 3})
	chan1.Close()

	chan2 := makeChannel[int]([]int{10, 20, 30})
	chan2.Close()

	mergeChan := makeChannel[int]([]int{100, 200, 300})
	mergeChan.Close()

	mergeChan.Merge(chan1, chan2)

	values := make([]int, 0, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		mergeChan.Range(func(value int) {
			values = append(values, value)
		})
		wg.Done()
	}()

	mergeChan.Write(111, 222)

	mergeChan.Close()
	wg.Wait()

	fmt.Print(values)
	// Output: [1 2 3 10 20 30 100 200 300 111 222]
}

func makeChannel[T any](arr []T) channels.Channel[T] {

	ch := channels.MakeChannel[T](len(arr))
	for _, num := range arr {
		ch <- num
	}

	return ch
}

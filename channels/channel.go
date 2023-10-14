package channels

import "sync"

type Channel[T any] chan T

func MakeChannel[T any](buff int) Channel[T] {
	if buff > 0 {
		return make(Channel[T], buff)
	} else {
		return make(Channel[T])
	}
}

func (ch Channel[T]) Check() bool {
	_, ok := <-ch
	return ok
}

func (ch Channel[T]) Safely() (T, bool) {
	val, ok := <-ch
	return val, ok
}

func (ch Channel[T]) Close() {
	close(ch)
}

func (ch Channel[T]) Len() int {
	return len(ch)
}

func (ch Channel[T]) Cap() int {
	return cap(ch)
}

func (ch Channel[T]) Write(values ...T) {
	for _, value := range values {
		ch <- value
	}
}

func (ch Channel[T]) Range(fn func(value T)) {
	for value := range ch {
		fn(value)
	}
}

func (ch *Channel[T]) Merge(channls ...Channel[T]) {

	var (
		wg   sync.WaitGroup
		size int
	)

	channls = append(channls, *ch)

	wg.Add(len(channls))

	for _, chann := range channls {
		size += chann.Len()
	}

	merged := MakeChannel[T](size)

	output := func(sc Channel[T]) {
		for sqr := range sc {
			merged.Write(sqr)
		}
		wg.Done()
	}

	for _, chann := range channls {
		output(chann)
	}

	wg.Wait()

	*ch = merged
}

//func (ch Channel[T]) Merge(channls ...Channel[T]) Channel[T] {
//
//	var (
//		wg   sync.WaitGroup
//		size int
//	)
//
//	channls = append(channls, ch)
//
//	wg.Add(len(channls))
//
//	for _, chann := range channls {
//		size += chann.Len()
//	}
//
//	merged := MakeChannel[T](size)
//
//	output := func(sc Channel[T]) {
//		for sqr := range sc {
//			merged.Write(sqr)
//		}
//		wg.Done()
//	}
//
//	for _, chann := range channls {
//		output(chann)
//	}
//
//	wg.Wait()
//
//	return merged
//}

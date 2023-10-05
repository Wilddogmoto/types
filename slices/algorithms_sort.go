package slices

import (
	"cmp"
	"slices"
)

type SortAlgorithm[E cmp.Ordered] func(arr Slice[E]) Slice[E]

func AnySort[E cmp.Ordered](compare func(a E, b E) int) SortAlgorithm[E] {

	return func(slice Slice[E]) Slice[E] {

		slices.SortFunc(slice, compare)

		return slice
	}
}

func QuickSort[E cmp.Ordered](array Slice[E]) Slice[E] {

	lenArray := array.Len()

	if lenArray < 2 {
		return array
	}

	pivot := array.Get(0)

	less := MakeSlice[E](0, lenArray)
	greater := MakeSlice[E](0, lenArray)

	for _, element := range array[1:] {

		if element <= pivot {
			less.Append(element)
		} else {
			greater.Append(element)
		}
	}

	lessSlice := QuickSort(less)
	greaterSlice := QuickSort(greater)

	out := MakeSlice[E](0, lessSlice.Len()+greaterSlice.Len())

	out.Append(lessSlice...)
	out.Append(pivot)
	out.Append(greaterSlice...)

	return out
}

/*
BubbleSort
Сортировка "Пузырьком"
Это очень простой алгоритм. Вам нужно сравнить каждый элемент массива со следующим элементом,
чтобы увидеть, больше ли он, если да, то вам нужно поменять их местами. Вы должны продолжать выполнять эту задачу,
пока больше нечего будет переставлять.
*/
func BubbleSort[E cmp.Ordered](array Slice[E]) Slice[E] {

	var isDone bool

	for !isDone {

		isDone = true
		i := 0

		for i < len(array)-1 {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isDone = false
			}
			i++
		}

	}

	return array

}

/*
InsertionSort
Сортировка вставками
Это алгоритм сортировки, в котором элементы входного массива поочередно выбираются и вставляются в отсортированную последовательность элементов.
Каждый новый элемент сравнивается с уже отсортированными элементами, и вставляется в нужное место в последовательности.
Этот процесс продолжается до тех пор, пока все элементы не будут отсортированы.
*/
func InsertionSort[E cmp.Ordered](array Slice[E]) Slice[E] {

	i := 1

	for i < len(array) {

		j := i

		for j >= 1 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]

			j--
		}

		i++
	}

	return array
}

/*
MergeSort
Сортировка слиянием
Это алгоритм сортировки, который упорядочивает элементы массива путем разделения его на две половины,
сортировки каждой половины отдельно, а затем слияния отсортированных половин в один отсортированный массив.
Алгоритм сортировки слиянием является эффективным и обычно используется для сортировки больших массивов.
Время выполнения сортировки слиянием в худшем, среднем и лучшем случае составляет O(n log n), где n - количество элементов в массиве.
*/
func MergeSort[E cmp.Ordered](array Slice[E]) Slice[E] {

	lenArray := array.Len()

	if lenArray == 1 {
		return array
	}

	fp := mergeSort[E](array[0 : lenArray/2])
	sp := mergeSort[E](array[lenArray/2:])

	return merge[E](fp, sp)
}

func mergeSort[E cmp.Ordered](array Slice[E]) Slice[E] {

	lenArray := array.Len()

	if lenArray == 1 {
		return array
	}

	fp := mergeSort[E](array[0 : lenArray/2])
	sp := mergeSort[E](array[lenArray/2:])

	return merge[E](fp, sp)
}

func merge[E cmp.Ordered](fp, sp Slice[E]) Slice[E] {

	//out := make(Slice[E], len(fp)+len(sp))

	lc := fp.Len() + sp.Len()

	out := MakeSlice[E](lc, lc)

	fpIndex := 0
	spIndex := 0
	nIndex := 0

	for fpIndex < fp.Len() && spIndex < sp.Len() {
		if fp[fpIndex] < sp[spIndex] {
			out[nIndex] = fp[fpIndex]
			fpIndex++
		} else if sp[spIndex] < fp[fpIndex] {
			out[nIndex] = sp[spIndex]
			spIndex++
		} else if fp[fpIndex] == sp[spIndex] {
			out[nIndex] = fp[fpIndex]
			fpIndex++
		}

		nIndex++
	}

	for fpIndex < fp.Len() {
		out[nIndex] = fp[fpIndex]

		fpIndex++
		nIndex++
	}

	for spIndex < sp.Len() {
		out[nIndex] = sp[spIndex]

		spIndex++
		nIndex++
	}

	return out
}

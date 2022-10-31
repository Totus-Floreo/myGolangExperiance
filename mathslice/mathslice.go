package mathslice

type Slice []Element

type Element int

// SumSlice — возвращает сумму элементов
func SumSlice(slice Slice) (res Element) {
	for _, s := range slice {
		res += s
	}
	return
}

// MapSlice — применяет функцию op к каждому элементу
func MapSlice(slice Slice, op func(Element) Element) {
	for i, s := range slice {
		slice[i] = op(s)
	}
}

// FoldSlice — сворачивает слайс.
func FoldSlice(slice Slice, op func(Element, Element) Element, init Element) (res Element) {
	res = op(init, slice[0])
	for i := 1; i < len(slice); i++ {
		res = op(res, slice[i])
	}
	return
}

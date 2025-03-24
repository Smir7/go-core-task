package main

func twoSlice(slice1, slice2 []string) []string {

	uniqueElements := map[string]bool{}

	for _, element := range slice2 {
		uniqueElements[element] = true
	}
	var result []string
	for _, element := range slice1 {
		if !uniqueElements[element] {
			result = append(result, element)
			delete(uniqueElements, element)
		}
	}
	return result

}

// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.
func main() {
	TestTwoSlice(twoSlice([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
}

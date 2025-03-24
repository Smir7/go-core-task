package main

func existIntersect(a, b []int) bool {
	uniqueElements := make(map[int]bool)
	for _, num := range a {
		uniqueElements[num] = true
	}
	for _, num := range b {
		if uniqueElements[num] {
			return true
		}
	}
	return false
}

func main() {

}

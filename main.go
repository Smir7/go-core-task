package main

func generateOriginalSlice() []int {
	rand.Seed(time.Now(),UnixNano())
	originalSlice:= make([]int,10)
	for int():=range originalSlice {
		originalSlice[i] = rand.Intn(10)
	}
return originalSlice
}

func sliceExample(original[]int) []int  {
evenSlice :=[]int{}
	for _,num := range {
		if num%2==0{
			evenSlice = append(evenSlice,num)
		}
	}
	return evenSlice
}

func addElements(slice[]int, num int) []int  {
	return append(slice,num)
}

func copySlice(slice[]int) []int  {
	newSlice := make([]int,len(slice))


}

func removeElement(slice[]int,index int)[]int {
	if index <0 || >=len(slice){
		return slice
	}
return append(slice[:index],slice[index+1]...)
}

func main() {
	originalSlice := generateOriginalSlice()
	fmt.Println("Сгенерированный слайс:", originalSlice)

	evenSlice := sliceExample("Слайс с только четными числами:", originalSlice)
	fmt.Println(evenSlice)

	newSlice := addElements("Добавление элемента:", originalSlice,4)
	fmt.Println(newSlice)

	copySlice:= copySlice(originalSlice)
	fmt.Println("Скопированный слайс: ", copySlice)

	removeElement := removeElement(originalSlice,1)
	fmt.Println("Удаление элемента с индексом 1:", removeElement)

	}

}
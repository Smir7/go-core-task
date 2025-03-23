package main
import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	output := sliceExample(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	num := 4
	expected := []int{1, 2, 3, 4}
	output := addElements(input, num)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	output := copySlice(input)
	output[0] = 10
	if input[0] == output[0] {
		t.Error("Copying failed, changes in the copy affect original.")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}
	index := 1
	expected := []int{1, 3, 4}
	output := removeElement(input, index)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}


	outputInvalid := removeElement(input, 10)
	if !reflect.DeepEqual(outputInvalid, input) {
		t.Errorf("Expected %v, but got %v for invalid index", input, outputInvalid)
	}
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func CreateValues() (string, string, string) {

	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	values := []string{
		fmt.Sprintf("%d", numDecimal),
		fmt.Sprintf("%d", numOctal),
		fmt.Sprintf("%d", numHexadecimal),
		fmt.Sprintf("%f", pi),
		fmt.Sprintf("%s", name),
		fmt.Sprintf("%t", isActive),
		fmt.Sprintf("%v", complexNum),
	}

	//var values []string
	//values = append(data, fmt.Sprintf("%d,%d,%d,%f,%s,%t,%v", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum))
	combinedString := strings.Join(values, ",")
	typeValues := fmt.Sprintf("Тип numDecimal: %s, Тип numOctal: %s, Тип numHexadecimal: %s, Тип pi: %s, Тип name: %s, Тип isActive: %s, Тип complexNum: %s",
		reflect.TypeOf(numDecimal), reflect.TypeOf(numOctal), reflect.TypeOf(numHexadecimal), reflect.TypeOf(pi), reflect.TypeOf(name), reflect.TypeOf(isActive), reflect.TypeOf(complexNum))

	runeSlice := []rune(combinedString)

	salt := "go-2024"
	midIndex := len(runeSlice) / 2
	saltedRunes := append(runeSlice[:midIndex], append([]rune(salt), runeSlice[midIndex:]...)...)

	hash := sha256.Sum256([]byte(string(saltedRunes)))
	hashedString := hex.EncodeToString(hash[:])

	return combinedString, hashedString, typeValues
}
func main() {
	combinedString, hashedString, typeValues := CreateValues()
	fmt.Println("Объединенная строка:", combinedString)
	fmt.Println("Хэш:", hashedString)
	fmt.Println("Типы переменных:", typeValues)
}

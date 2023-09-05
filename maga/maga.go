package maga

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	elems, err := readElems()
	if err != nil {
		log.Fatal(err)
	}
	result := countDigitsInArray(elems)
	fmt.Printf("количесвто цифр в массиве, %v\n", result)
}

// функция считывания из консоли массива элементов
func readElems() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("недостаточное количество входных данных")
	}
	elemsLen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return nil, err
	}
	if elemsLen < 0 {
		return nil, errors.New("длина массива не может быть отрицательной")
	}
	elems := os.Args[2:]
	if elemsLen != len(elems) {
		return nil, errors.New("длина массива не соответсвтует количеству элементов")
	}

	return elems, nil
}

// функция подсчета цифр в массиве
func countDigitsInArray(arr []string) int {
	count := 0
	for _, str := range arr {
		for _, char := range str {
			if char >= '0' && char <= '9' {
				count++
			}
		}
	}
	return count
}


package myscanner

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ScanFloat(filename string) []float64{
	file, err := os.Open(filename)
	errorHandler(err)

	scanner := bufio.NewScanner(file)	
	var data []float64 
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		errorHandler(err)
		data = append(data, num)
	}

	errorHandler(file.Close())
	errorHandler(scanner.Err())
	return data
}

func ScanWords(filename string) []string{
	file, err := os.Open(filename)
	errorHandler(err)

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	return words
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}